CREATE TABLE clinic (
                         id serial NOT NULL,
                         title varchar NOT NULL,
                         created  timestamptz NOT NULL DEFAULT NOW(),
                         updated  timestamptz NULL,
                         CONSTRAINT clinic_pk PRIMARY KEY (id)
);

CREATE TABLE patient (
                          id serial NOT NULL,
                          status_code varchar NOT NULL,
                          created  timestamptz NOT NULL DEFAULT NOW(),
                          updated  timestamptz NULL,
                          CONSTRAINT patient_pk PRIMARY KEY (id)
);

CREATE TABLE employee (
                           id uuid NOT NULL,
                           guid uuid NOT NULL UNIQUE,
                           first_name varchar NOT NULL,
                           last_name varchar NOT NULL,
                           patronymic varchar NULL,
                           position_code varchar NOT NULL,
                           photo_guid uuid NULL,
                           created timestamptz NOT NULL DEFAULT NOW(),
                           updated timestamptz NULL,
                           CONSTRAINT employee_pk PRIMARY KEY (id)
);

CREATE TABLE appointment_program (
                       id serial NOT NULL,
                       doctor_id int NOT NULL,
                       type_code varchar NOT NULL,
                       status_code varchar NOT NULL,
                       begin timestamptz NOT NULL,
                       "end" timestamptz NULL,
                       comment text NULL,
                       periodicity varchar NOT NULL,
                       created timestamptz NOT NULL DEFAULT NOW(),
                       updated timestamptz NULL,
                       CONSTRAINT appointment_program_pk PRIMARY KEY (id)
);

CREATE TABLE appointment (
                           id serial NOT NULL,
                           created  timestamptz NOT NULL DEFAULT NOW(),
                           updated  timestamptz NULL,
                           doctor_id int NULL,
                           clinic_id int NOT NULL REFERENCES clinic(id) ON DELETE CASCADE,
                           program_id int NOT NULL REFERENCES appointment_program(id) ON DELETE CASCADE,
                           patient_id int NOT NULL REFERENCES patient(id) ON DELETE CASCADE,
                           type_code varchar NOT NULL,
                           status_code varchar NOT NULL,
                           planned timestamptz NOT NULL,
                           performed timestamptz NULL,
                           duration timestamptz NULL,
                           comment text NULL,
                           CONSTRAINT appointment_pk PRIMARY KEY (id)
);

CREATE TABLE appointment_param (
                         id serial NOT NULL,
                         created  timestamptz NOT NULL DEFAULT NOW(),
                         updated  timestamptz NULL,
                         appointment_id int NOT NULL UNIQUE REFERENCES appointment(id) ON DELETE CASCADE,
                         type_code varchar NOT NULL,
                         value varchar NOT NULL,
                         CONSTRAINT appointment_param_pk PRIMARY KEY (id)
);

CREATE TABLE confirm_code (
                            id serial NOT NULL,
                            patient_id int NOT NULL REFERENCES patient(id) ON DELETE CASCADE,
                            channel_id int NOT NULL,

                            active BOOLEAN NOT NULL,
                            type_code varchar NOT NULL,
                            code varchar NOT NULL,
                            created timestamptz NOT NULL DEFAULT NOW(),
                            updated timestamptz NULL,
                            CONSTRAINT confirm_code_pk PRIMARY KEY (id)

);

CREATE TABLE implementation (
                           id serial NOT NULL,
                           appointment_id int NOT NULL REFERENCES appointment(id) ON DELETE CASCADE,
                           status_code varchar NOT NULL check (
                                       status_code = 'PLANNED'
                                   or status_code = 'ONLINE'
                                   or status_code = 'FINISHED'
                               ),
                           performed timestamptz NULL,
                           created timestamptz NOT NULL DEFAULT NOW(),
                           updated timestamptz NULL,
                           CONSTRAINT implementation_pk PRIMARY KEY (id)
);

CREATE TABLE implementation_param (
                             id serial NOT NULL,
                             implementation_id int NOT NULL REFERENCES implementation(id) ON DELETE CASCADE,
                             created timestamptz NOT NULL DEFAULT NOW(),
                             updated timestamptz NULL,
                             type_code varchar NOT NULL,
                             value varchar NULL,
                             CONSTRAINT implementation_param_pk PRIMARY KEY (id)
);

CREATE TABLE invite (
                              id serial NOT NULL,
                              patient_id int NOT NULL REFERENCES patient(id) ON DELETE CASCADE,
                              channel_id int NOT NULL,
                              status_code varchar NOT NULL,
                              created timestamptz NOT NULL DEFAULT NOW(),
                              updated timestamptz NULL,
                              CONSTRAINT invite_pk PRIMARY KEY (id)

);
