job "datarouter-${CI_REGISTRY_VER}" {
  datacenters = ["MS001"]
  region = "Moscow"
  type = "service"
  constraint {
    attribute = "${node.class}"
    value = "private"
  }
  update {
    max_parallel = 1
    min_healthy_time = "2m"
    healthy_deadline = "6m"
    auto_revert = false
    canary = 0
  }
  group "datarouter" {
    count = 1
    restart {
      attempts = 4
      interval = "5m"
      delay = "1m"
      mode = "delay"
    }

    task "config-db" {
      vault {
        policies = ["kv"]
        change_mode = "restart"
      }
      driver = "docker"
      template {
data = <<EOH
MONGO_INITDB_DATABASE="{{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_DATABASE}}{{end}}"
EOH
        destination = "secrets/file.env"
        change_mode = "restart"
        splay = "15s"
        env = true
      }
      template {
data = <<EOH
{{with secret `secret/datarouter/mongodb_dump`}}{{.Data.DUMP}}{{end}}
EOH
        destination = "local/init.js"
        change_mode = "restart"
        splay = "15s"
      }
      config {
        hostname = "config-db"
        force_pull = false
        image = "mongo"
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = ["consul"]
        volumes = [ "/opt/maximus/platform/nomad/data/alloc/${NOMAD_ALLOC_ID}/${NOMAD_TASK_NAME}/local/init.js:/docker-entrypoint-initdb.d/init.js:ro" ]
        port_map { mongodb = 27017 }
      }
      env {
        MONGO_DATA_DIR = "/data/db"
        MONGO_LOG_DIR = "/dev/null"
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
        name = "datarouter-${CI_NOMAD_VER}"
        tags = ["config-db"]
        port = "mongodb"
        check {
          port = "mongodb"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 10
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
INFLUXDB_DB="{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_DB}}{{end}}"
INFLUXDB_USER="{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_USER}}{{end}}"
INFLUXDB_USER_PASSWORD="{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_USER_PASSWORD}}{{end}}"
EOH
        destination = "secrets/file.env"
        change_mode = "restart"
        splay = "15s"
        env = true
      }
      config {
        hostname = "event-db"
        force_pull = false
        image = "influxdb:latest"
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = ["consul"]
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
        name = "datarouter-${CI_NOMAD_VER}"
        tags = ["event-db"]
        port = "influxdb"
        check {
          port = "influxdb"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 10
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }

    task "memcached" {
      driver = "docker"
      config {
        hostname = "memcached"
        image = "memcached:latest"
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = ["consul"]
        port_map { memcached = 11211 }
      }
      resources {
        cpu = 256
        memory = 128
        network {
          mbits = 10
          port "memcached" {}
        }
      }
      service {
        name = "datarouter-${CI_NOMAD_VER}"
        tags = ["memcached"]
        port = "memcached"
        check {
          port = "memcached"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 10
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }

    task "broker" {
      vault {
        policies = ["kv"]
        change_mode = "restart"
      }
      driver = "docker"
      template {
data = <<EOH
allow_anonymous = off
allow_register_during_netsplit = off
allow_publish_during_netsplit = off
allow_subscribe_during_netsplit = off
allow_unsubscribe_during_netsplit = off
allow_multiple_sessions = on
queue_deliver_mode = fanout
retry_interval = 5
## max_client_id_size = 100
## persistent_client_expiration = 1w
max_inflight_messages = 1
max_online_messages = -1
max_offline_messages = 1000
max_message_rate = 0
max_message_size = 0
upgrade_outgoing_qos = off
listener.max_connections = 1000000
listener.nr_of_acceptors = 10
listener.tcp.default = 0.0.0.0:1883
#listener.ssl.default = 0.0.0.0:8883
listener.ws.default = 0.0.0.0:8080
listener.vmq.clustering = 0.0.0.0:44053
listener.http.default = 0.0.0.0:8888
listener.mountpoint = off

systree_enabled = on
systree_interval = 20000

graphite_enabled = off
graphite_host = localhost
graphite_port = 2003
graphite_interval = 20000

shared_subscription_policy = random

plugins.vmq_passwd = off
plugins.vmq_acl = off
plugins.vmq_diversity = on
plugins.vmq_webhooks = off
plugins.vmq_bridge = off
vmq_acl.acl_file = /etc/vernemq/vmq.acl
vmq_acl.acl_reload_interval = 10
vmq_passwd.password_file = /etc/vernemq/vmq.passwd
vmq_passwd.password_reload_interval = 10
vmq_diversity.script_dir = /opt/vernemq/share/lua

vmq_diversity.auth_mongodb.enabled = on
vmq_diversity.mongodb.host = {{ range service "config-db.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}
vmq_diversity.mongodb.port = {{ range service "config-db.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}}
#vmq_diversity.mongodb.login = {{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_USER}}{{end}}
#vmq_diversity.mongodb.password = {{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_USER_PASSWORD}}{{end}}
vmq_diversity.mongodb.database = {{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_DATABASE}}{{end}}

vmq_diversity.memcache.host = {{ range service "memcached.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}
vmq_diversity.memcache.port = {{ range service "memcached.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}}

log.console = console
log.console.level = info
log.console.file = /opt/vernemq/log/console.log
log.error.file = /opt/vernemq/log/error.log
log.syslog = off
log.crash = on
log.crash.file = /opt/vernemq/log/crash.log
log.crash.maximum_message_size = 64KB
log.crash.size = 10MB
log.crash.rotation = $D0
log.crash.rotation.keep = 5

nodename = broker@127.0.0.1
distributed_cookie = vmq

erlang.async_threads = 64
erlang.max_ports = 262144
leveldb.maximum_memory.percent = 70
EOH
        destination = "local/vernemq.conf"
        change_mode = "restart"
        splay = "15s"
      }
      config {
        hostname = "broker"
        force_pull = true
        image = "${REGISTRY_BROKER_URL}"
        auth {
          username = "${REGISTRY_USERNAME}"
          password = "${REGISTRY_PASSWORD}"
          server_address = "${REGISTRY_HOST}"
        }
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = ["consul"]
        volumes = [ "/opt/maximus/platform/nomad/data/alloc/${NOMAD_ALLOC_ID}/${NOMAD_TASK_NAME}/local/vernemq.conf:/opt/vernemq/etc/vernemq.conf:ro" ]
        port_map {
          vernemq = 1883
          metrics = 8888
        }
      }
      resources {
        cpu = 1024
        memory = 256
        network {
          mbits = 10
          port "vernemq" {}
          port "metrics" {}
        }
      }
      service {
        name = "datarouter-${CI_NOMAD_VER}"
        tags = ["broker"]
        port = "vernemq"
        check {
          port = "vernemq"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 10
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
      service {
        name = "datarouter-${CI_NOMAD_VER}"
        tags = [ "task=broker", "metrics" ]
        port = "metrics"
        check {
          port = "metrics"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 10
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }

    task "api" {
      vault {
        policies = ["kv"]
        change_mode = "restart"
      }
      driver = "docker"
      template {
data = <<EOH
{
 "http": {
  "host": "0.0.0.0",
  "port": 80
 },
 "configDB": {
  "host": "{{ range service "config-db.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}",
  "port": {{ range service "config-db.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}},
  "database": "{{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_DATABASE}}{{end}}"
 },
 "eventDB": {
  "host": "{{ range service "event-db.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}",
  "port": {{ range service "event-db.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}},
  "protocol": "http",
  "database": "{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_DB}}{{end}}",
  "login": "{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_USER}}{{end}}",
  "password": "{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_USER_PASSWORD}}{{end}}"
 },
 "sentryDSN": "{{with secret `secret/datarouter/sentryDSN`}}{{.Data.value}}{{end}}",
 "logging": {
  "output": "STDOUT",
  "level": "DEBUG",
  "format": "TEXT"
 }
}
EOH
        destination = "local/config.json"
        change_mode = "restart"
        splay = "15s"
      }
      config {
        hostname = "api"
        force_pull = true
        image = "${REGISTRY_API_URL}:${CI_REGISTRY_VER}"
        auth {
          username = "${REGISTRY_USERNAME}"
          password = "${REGISTRY_PASSWORD}"
          server_address = "${REGISTRY_HOST}"
        }
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = ["consul"]
        volumes = [ "/opt/maximus/platform/nomad/data/alloc/${NOMAD_ALLOC_ID}/${NOMAD_TASK_NAME}/local/config.json:/opt/config.json:ro" ]
        port_map { http = 80 }
        args = [ "-c", "/opt/config.json" ]
      }
      resources {
        cpu = 256
        memory = 128
        network {
          mbits = 10
          port "http" {}
        }
      }
      service {
        name = "datarouter-${CI_NOMAD_VER}"
        tags = [ "api", "traefik.enable=true", "traefik.frontend.rule=Host:test.maximus.lan;PathPrefixStrip:/datarouter/${CI_REGISTRY_VER}" ]
        port = "http"
        check {
          port = "http"
          type = "tcp"
          interval = "15s"
          timeout = "5s"
          check_restart {
            limit = 10
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }

    task "proxy" {
      vault {
        policies = ["kv"]
        change_mode = "restart"
      }
      driver = "docker"
      template {
data = <<EOH
{
 "configDB": {
  "host": "{{ range service "config-db.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}",
  "port": {{ range service "config-db.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}},
  "database": "{{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_DATABASE}}{{end}}"
 },
 "eventDB": {
  "host": "{{ range service "event-db.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}",
  "port": {{ range service "event-db.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}},
  "protocol": "http",
  "database": "{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_DB}}{{end}}",
  "login": "{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_USER}}{{end}}",
  "password": "{{with secret `secret/datarouter/influxdb`}}{{.Data.INFLUXDB_USER_PASSWORD}}{{end}}"
 },
 "sentryDSN": "{{with secret `secret/datarouter/sentryDSN`}}{{.Data.value}}{{end}}",
 "logging": {
  "output": "STDOUT",
  "level": "DEBUG",
  "format": "TEXT"
 },
"prometheus": {
  "port": 9100,
  "path": "/metrics"
 },
 "mq": {
  "host": "{{ range service "broker.datarouter-${CI_NOMAD_VER}" }}{{ .Address }}{{end}}",
  "port": {{ range service "broker.datarouter-${CI_NOMAD_VER}" }}{{ .Port }}{{end}},
  "pubClientID": "publisher_dataRouterID",
  "subClientID": "subscriber_dataRouterID",
  "login": "{{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_USER}}{{end}}",
  "password": "{{with secret `secret/datarouter/mongodb`}}{{.Data.MONGO_USER_PASSWORD}}{{end}}"
 }
}
EOH
        destination = "local/config.json"
        change_mode = "restart"
        splay = "15s"
      }
      config {
        hostname = "proxy"
        force_pull = true
        image = "${REGISTRY_PROXY_URL}:${CI_REGISTRY_VER}"
        auth {
          username = "${REGISTRY_USERNAME}"
          password = "${REGISTRY_PASSWORD}"
          server_address = "${REGISTRY_HOST}"
        }
        dns_servers = [ "${attr.unique.network.ip-address}" ]
        dns_search_domains = ["consul"]
        volumes = [ "/opt/maximus/platform/nomad/data/alloc/${NOMAD_ALLOC_ID}/${NOMAD_TASK_NAME}/local/config.json:/opt/config.json:ro" ]
        port_map { metrics  = 9100 }
        args = [ "-c", "/opt/config.json" ]
      }
      resources {
        cpu = 256
        memory = 128
        network {
          mbits = 10
          port "metrics" {}
        }
      }
      service {
        name = "datarouter-${CI_NOMAD_VER}"
        tags = ["task=proxy", "metrics"]
        port = "metrics"
        check {
          port = "metrics"
          type = "tcp"
          interval = "1m"
          timeout = "30s"
          check_restart {
            limit = 10
            grace = "90s"
            ignore_warnings = true
          }
        }
      }
    }
  }
}
