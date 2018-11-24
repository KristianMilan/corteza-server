// Code generated by statik. DO NOT EDIT.

// Package statik contains static assets.
package mysql

import (
	"github.com/rakyll/statik/fs"
)

func Data() string {
	return "PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1a\x00	\x0020180704080000.base.up.sqlUT\x05\x00\x01\x80Cm8-- all known organisations (crust instances) and our relation towards them\nCREATE TABLE organisations (\n  id               BIGINT UNSIGNED NOT NULL,\n  fqn              TEXT            NOT NULL, -- fully qualified name of the organisation\n  name             TEXT            NOT NULL, -- display name of the organisation\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  archived_at      DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- organisation soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE settings (\n  name  VARCHAR(200) NOT NULL   COMMENT 'Unique set of setting keys',\n  value TEXT                    COMMENT 'Setting value',\n\n  PRIMARY KEY (name)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\n-- Keeps all known users, home and external organisation\n--   changes are stored in audit log\nCREATE TABLE users (\n  id               BIGINT UNSIGNED NOT NULL,\n  email            TEXT            NOT NULL,\n  username         TEXT            NOT NULL,\n  password         TEXT            NOT NULL,\n  name             TEXT            NOT NULL,\n  handle           TEXT            NOT NULL,\n  meta             JSON            NOT NULL,\n  satosa_id        CHAR(36)            NULL,\n\n  rel_organisation BIGINT UNSIGNED NOT NULL,\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  suspended_at     DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- user soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE UNIQUE INDEX uid_satosa ON users (satosa_id);\n\n-- Keeps all known teams\nCREATE TABLE teams (\n  id               BIGINT UNSIGNED NOT NULL,\n  name             TEXT            NOT NULL, -- display name of the team\n  handle           TEXT            NOT NULL, -- team handle string\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  archived_at      DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- team soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\n-- Keeps team memberships\nCREATE TABLE team_members (\n  rel_team         BIGINT UNSIGNED NOT NULL REFERENCES organisation(id),\n  rel_user         BIGINT UNSIGNED NOT NULL,\n\n  PRIMARY KEY (rel_team, rel_user)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\nPK\x07\x08\xedzU\x8am	\x00\x00m	\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00.\x00	\x0020181124181811.rename_and_prefix_tables.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE teams RENAME TO sys_team;\nALTER TABLE organisations RENAME TO sys_organisation;\nALTER TABLE team_members RENAME TO sys_team_member;\nALTER TABLE users RENAME TO sys_user;PK\x07\x08\xf2\xc4\x87\xe8\xb5\x00\x00\x00\xb5\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00migrations.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS `migrations` (\n `project` varchar(16) NOT NULL COMMENT 'sam, crm, ...',\n `filename` varchar(255) NOT NULL COMMENT 'yyyymmddHHMMSS.sql',\n `statement_index` int(11) NOT NULL COMMENT 'Statement number from SQL file',\n `status` TEXT NOT NULL COMMENT 'ok or full error message',\n PRIMARY KEY (`project`,`filename`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nPK\x07\x08\x0d\xa5T2x\x01\x00\x00x\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x06\x00	\x00new.shUT\x05\x00\x01\x80Cm8#!/bin/bash\ntouch $(date +%Y%m%d%H%M%S).up.sqlPK\x07\x08s\xd4N*.\x00\x00\x00.\x00\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xedzU\x8am	\x00\x00m	\x00\x00\x1a\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x0020180704080000.base.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xf2\xc4\x87\xe8\xb5\x00\x00\x00\xb5\x00\x00\x00.\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xbe	\x00\x0020181124181811.rename_and_prefix_tables.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x0d\xa5T2x\x01\x00\x00x\x01\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xd8\n\x00\x00migrations.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(s\xd4N*.\x00\x00\x00.\x00\x00\x00\x06\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xed\x81\x95\x0c\x00\x00new.shUT\x05\x00\x01\x80Cm8PK\x05\x06\x00\x00\x00\x00\x04\x00\x04\x008\x01\x00\x00\x00\x0d\x00\x00\x00\x00"
}

func init() {
	fs.Register(Data())
}
