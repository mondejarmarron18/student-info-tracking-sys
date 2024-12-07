
CREATE TABLE IF NOT EXISTS address
(
  id         uuid      NOT NULL DEFAULT gen_random_uuid(),
  student_id uuid      NOT NULL,
  city       varchar   NOT NULL,
  country    varchar   NOT NULL,
  barangay   varchar   NOT NULL,
  street     varchar   NOT NULL,
  type       varchar   NOT NULL DEFAULT 'permanent',
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN address.type IS 'permanent | current';

CREATE TABLE IF NOT EXISTS course
(
  id          uuid      NOT NULL DEFAULT gen_random_uuid(),
  name        varchar   NOT NULL UNIQUE,
  description text      NOT NULL,
  created_at  timestamp NOT NULL DEFAULT NOW(),
  updated_at  timestamp NOT NULL DEFAULT NOW(),
  deleted_at  timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS course_level
(
  id         uuid      NOT NULL DEFAULT gen_random_uuid(),
  course_id  uuid      NOT NULL,
  student_id uuid      NOT NULL,
  level      int       NOT NULL DEFAULT 1,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN course_level.student_id IS 'user_profile_id';

CREATE TABLE IF NOT EXISTS course_specialization
(
  id              uuid      NOT NULL DEFAULT gen_random_uuid(),
  course_level_id uuid      NOT NULL,
  name            varchar   NOT NULL UNIQUE,
  description     text      NOT NULL,
  created_at      timestamp NOT NULL DEFAULT NOW(),
  updated_at      timestamp NOT NULL DEFAULT NOW(),
  deleted_at      timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS permission
(
  id          uuid      NOT NULL DEFAULT gen_random_uuid(),
  name        varchar   NOT NULL UNIQUE,
  description text     ,
  created_at  timestamp NOT NULL DEFAULT NOW(),
  updated_at  timestamp NOT NULL DEFAULT NOW(),
  deleted_at  timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN permission.name IS 'view:students | view:student | delete:student ...';

CREATE TABLE IF NOT EXISTS role
(
  id          uuid      NOT NULL DEFAULT gen_random_uuid(),
  name        varchar   NOT NULL UNIQUE,
  description text     ,
  created_at  timestamp NOT NULL DEFAULT NOW(),
  updated_at  timestamp NOT NULL DEFAULT NOW(),
  deleted_at  timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN role.name IS 'student | admin | staff ...';

CREATE TABLE IF NOT EXISTS role_permission
(
  id            uuid      NOT NULL DEFAULT gen_random_uuid(),
  permission_id uuid      NOT NULL,
  role_id       uuid      NOT NULL,
  created_at    timestamp NOT NULL DEFAULT NOW(),
  updated_at    timestamp NOT NULL DEFAULT NOW(),
  deleted_at    timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS student_guardian
(
  id           uuid      NOT NULL DEFAULT gen_random_uuid(),
  student_id   uuid      NOT NULL,
  first_name   varchar   NOT NULL,
  middle_name  varchar  ,
  last_name    varchar   NOT NULL,
  phone_number varchar   NOT NULL UNIQUE,
  created_at   timestamp NOT NULL DEFAULT NOW(),
  updated_at   timestamp NOT NULL DEFAULT NOW(),
  deleted_at   timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN student_guardian.student_id IS 'user_profile_id';

CREATE TABLE IF NOT EXISTS update_approval
(
  id                 uuid      NOT NULL DEFAULT gen_random_uuid(),
  reviewed_by        uuid      NOT NULL,
  table_name         varchar   NOT NULL,
  field_name         varchar   NOT NULL,
  prev_value         varchar   NOT NULL,
  new_value          varchar   NOT NULL,
  review_status      int       DEFAULT 2,
  review_description text     ,
  created_at         timestamp NOT NULL,
  deleted_at         timestamp,
  reviewed_at        timestamp,
  cancelled_at       timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN update_approval.table_name IS 'user_profile | adress ...';

COMMENT ON COLUMN update_approval.field_name IS 'name | barangay | region ...';

COMMENT ON COLUMN update_approval.review_status IS '0 = rejected | 1 = approved | 2 = pending';

CREATE TABLE IF NOT EXISTS user_account
(
  id          uuid      NOT NULL DEFAULT gen_random_uuid(),
  role_id     uuid      NOT NULL,
  email       varchar   NOT NULL UNIQUE,
  password    varchar   NOT NULL,
  verified_at varchar  ,
  deleted_at  timestamp,
  created_at  timestamp DEFAULT NOW(),
  updated_at  timestamp DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_profile
(
  id              uuid      NOT NULL DEFAULT gen_random_uuid(),
  user_account_id uuid      NOT NULL,
  id_number       varchar   NOT NULL UNIQUE,
  first_name      varchar   NOT NULL,
  middle_name     varchar  ,
  last_name       varchar   NOT NULL,
  phone_number    varchar   NOT NULL UNIQUE,
  contact_methods varchar   DEFAULT 'email',
  created_at      timestamp NOT NULL DEFAULT NOW(),
  updated_at      timestamp NOT NULL DEFAULT NOW(),
  deleted_at      timestamp,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN user_profile.contact_methods IS 'email, sms, phone call';

ALTER TABLE user_profile
  ADD CONSTRAINT FK_user_account_TO_user_profile
    FOREIGN KEY (user_account_id)
    REFERENCES user_account (id);

ALTER TABLE address
  ADD CONSTRAINT FK_user_profile_TO_address
    FOREIGN KEY (student_id)
    REFERENCES user_profile (id);

ALTER TABLE student_guardian
  ADD CONSTRAINT FK_user_profile_TO_student_guardian
    FOREIGN KEY (student_id)
    REFERENCES user_profile (id);

ALTER TABLE role_permission
  ADD CONSTRAINT FK_permission_TO_role_permission
    FOREIGN KEY (permission_id)
    REFERENCES permission (id);

ALTER TABLE role_permission
  ADD CONSTRAINT FK_role_TO_role_permission
    FOREIGN KEY (role_id)
    REFERENCES role (id);

ALTER TABLE course_level
  ADD CONSTRAINT FK_course_TO_course_level
    FOREIGN KEY (course_id)
    REFERENCES course (id);

ALTER TABLE course_specialization
  ADD CONSTRAINT FK_course_level_TO_course_specialization
    FOREIGN KEY (course_level_id)
    REFERENCES course_level (id);

ALTER TABLE course_level
  ADD CONSTRAINT FK_user_profile_TO_course_level
    FOREIGN KEY (student_id)
    REFERENCES user_profile (id);

ALTER TABLE update_approval
  ADD CONSTRAINT FK_address_TO_update_approval
    FOREIGN KEY (reviewed_by)
    REFERENCES address (id);

ALTER TABLE user_account
  ADD CONSTRAINT FK_role_TO_user_account
    FOREIGN KEY (role_id)
    REFERENCES role (id);
