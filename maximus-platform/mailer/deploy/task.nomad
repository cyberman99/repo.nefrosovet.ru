job "mailer-${CI_REGISTRY_VER}" {
  datacenters = [ "MS001" ]
  type = "service"
  update {
    max_parallel = 1
    min_healthy_time = "2m"
    healthy_deadline = "6m"
    auto_revert = false
    canary = 0
  }
  group "mailer" {
    count = 1
    restart {
      attempts = 4
      interval = "5m"
      delay = "1m"
      mode = "delay"
    }
    task "config-db" {
      driver = "docker"
      config {
        hostname = "config-db"
        force_pull = false
        image = "mongo"
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = [ "consul" ]
        volumes = [ "/opt/maximus/platform/nomad/data/alloc/${NOMAD_ALLOC_ID}/${NOMAD_TASK_NAME}${NOMAD_TASK_DIR}:/data/db:rw" ]
        port_map { mongodb = 27017 }
      }
      env {
        MONGO_DATA_DIR="/data/db"
        MONGO_LOG_DIR="/dev/null"
      }
      resources {
        cpu = 256
        memory = 128
        network {
          mbits = 10
          port "mongodb" {}
        }
      }
      service {
        name = "mailer-${CI_NOMAD_VER}"
        tags = [ "config-db" ]
        port = "mongodb"
        check {
          port = "mongodb"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 3
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }
    task "event-db" {
      vault {
        policies = ["kv"]
        change_mode = "restart"
      }
      driver = "docker"
      template {
data = <<EOH
INFLUXDB_DB="{{with secret `secret/mailer/influxdb`}}{{.Data.INFLUXDB_DB}}{{end}}"
INFLUXDB_USER="{{with secret `secret/mailer/influxdb`}}{{.Data.INFLUXDB_USER}}{{end}}"
INFLUXDB_USER_PASSWORD="{{with secret `secret/mailer/influxdb`}}{{.Data.INFLUXDB_USER_PASSWORD}}{{end}}"
EOH
        destination = "secrets/file.env"
        change_mode = "restart"
        splay = "15s"
        env = true
      }
      config {
        hostname = "event-db"
        force_pull = false
        image = "influxdb"
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = [ "consul" ]
        volumes = [ "/opt/maximus/platform/nomad/data/alloc/${NOMAD_ALLOC_ID}/${NOMAD_TASK_NAME}${NOMAD_TASK_DIR}:/var/lib/influxdb:rw" ]
        port_map { influxdb = 8086 }
      }
      resources {
        cpu = 256
        memory = 128
        network {
          mbits = 10
          port "influxdb" {}
        }
      }
      service {
        name = "mailer-${CI_NOMAD_VER}"
        tags = [ "event-db" ]
        port = "influxdb"
        check {
          port = "influxdb"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 3
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }
    task "app" {
      vault {
        policies = ["kv"]
        change_mode = "restart"
      }
      driver = "docker"
      template {
data = <<EOH
MAILER_EVENTDB_DATABASE="{{with secret `secret/mailer/influxdb`}}{{.Data.INFLUXDB_DB}}{{end}}"
MAILER_EVENTDB_LOGIN="{{with secret `secret/mailer/influxdb`}}{{.Data.INFLUXDB_USER}}{{end}}"
MAILER_EVENTDB_PASSWORD="{{with secret `secret/mailer/influxdb`}}{{.Data.INFLUXDB_USER_PASSWORD}}{{end}}"
MAILER_MASTERTOKEN="{{with secret `secret/mailer/app`}}{{.Data.MAILER_MASTERTOKEN}}{{end}}"
MAILER_SENTRY_DSN="{{with secret `secret/mailer/sentryDSN`}}{{.Data.value}}{{end}}"
MAILER_CONFIGDB_DATABASE="{{with secret `secret/mailer/mongodb`}}{{.Data.MONGODB_DB}}{{end}}"
MAILER_CONFIGDB_HOST = "{{ range service "config-db.mailer-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}"
MAILER_CONFIGDB_PORT = "{{ range service "config-db.mailer-${CI_NOMAD_VER}" }}{{ .Port }}{{end}}"
MAILER_EVENTDB_HOST = "{{ range service "event-db.mailer-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}"
MAILER_EVENTDB_PORT = "{{ range service "event-db.mailer-${CI_NOMAD_VER}" }}{{ .Port }}{{end}}"
EOH
        destination = "secrets/file.env"
        change_mode = "restart"
        splay = "15s"
        env = true
      }
      config {
        hostname = "app"
        force_pull = true
        image = "${REGISTRY_APP_URL}:${CI_REGISTRY_VER}"
        force_pull = true
        auth {
          username = "${REGISTRY_USERNAME}"
          password = "${REGISTRY_PASSWORD}"
          server_address = "${REGISTRY_HOST}"
        }
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = [ "consul" ]
        port_map {
          app = 80
          metrics = 9090
        }
      }
      env {
        MAILER_HTTP_HOST = "0.0.0.0"
        MAILER_HTTP_PORT = "80"
        MAILER_PROMETHEUS_PORT = "9090"
        MAILER_LOGGING_LEVEL = "INFO"
        MAILER_BOTPROXY_HTTP_HOST = "botproxy.diacare-soft.ru"
		MAILER_BOTPROXY_MQ_HOST = "botproxy.diacare-soft.ru"
		MAILER_BOTPROXY_MQ_PORT = "8883"
      }
      resources {
        cpu = 256
        memory = 128
        network {
          mbits = 100
          port "app" {}
          port "metrics" {}
        }
      }
      service {
        name = "mailer-${CI_NOMAD_VER}"
        tags = [ "mailer-${CI_REGISTRY_VER}", "traefik.enable=true", "traefik.frontend.rule=Host:test.maximus.lan;PathPrefixStrip:/mailer/${CI_REGISTRY_VER}" ]
        port = "app"
        check {
          port = "app"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 3
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
      service {
        name = "mailer-${CI_NOMAD_VER}"
        tags = [ "metrics" ]
        port = "metrics"
        check {
          port = "metrics"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 3
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }
  }
}
