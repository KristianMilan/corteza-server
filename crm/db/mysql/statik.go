// Code generated by statik. DO NOT EDIT.

package mysql

import (
	"github.com/rakyll/statik/fs"
)

func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00\x01\xa9HM\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1a\x00	\x0020180704080000.base.up.sqlUT\x05\x00\x013\xc7\xbb[CREATE TABLE `crm_content` (\n `id` bigint(20) unsigned NOT NULL,\n `module_id` bigint(20) unsigned NOT NULL,\n `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,\n `updated_at` datetime DEFAULT NULL,\n `deleted_at` datetime DEFAULT NULL,\n PRIMARY KEY (`id`,`module_id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE `crm_content_column` (\n `content_id` bigint(20) NOT NULL,\n `column_name` varchar(255) NOT NULL,\n `column_value` text NOT NULL,\n PRIMARY KEY (`content_id`,`column_name`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE `crm_field` (\n `field_type` varchar(16) NOT NULL COMMENT 'Short field type (string, boolean,...)',\n `field_name` varchar(255) NOT NULL COMMENT 'Description of field contents',\n `field_template` varchar(255) NOT NULL COMMENT 'HTML template file for field',\n PRIMARY KEY (`field_type`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE `crm_module` (\n `id` bigint(20) unsigned NOT NULL,\n `name` varchar(64) NOT NULL COMMENT 'The name of the module',\n `json` json NOT NULL COMMENT 'List of field definitions for the module',\n `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,\n `updated_at` datetime DEFAULT NULL,\n `deleted_at` datetime DEFAULT NULL,\n PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE `crm_module_form` (\n `module_id` bigint(20) unsigned NOT NULL,\n `place` tinyint(3) unsigned NOT NULL,\n `kind` varchar(64) NOT NULL COMMENT 'The type of the form input field',\n `name` varchar(64) NOT NULL COMMENT 'The name of the field in the form',\n `label` varchar(255) NOT NULL COMMENT 'The label of the form input',\n `help_text` text NOT NULL COMMENT 'Help text',\n `default_value` text NOT NULL COMMENT 'Default value',\n `max_length` int(10) unsigned NOT NULL COMMENT 'Maximum input length',\n `is_private` tinyint(1) NOT NULL COMMENT 'Contains personal/sensitive data?',\n PRIMARY KEY (`module_id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nPK\x07\x08E\x9b\x99\xdc{\x07\x00\x00{\x07\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x01\xa9HM\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00%\x00	\x0020180704080001.crm_fields-data.up.sqlUT\x05\x00\x013\xc7\xbb[INSERT INTO `crm_field` VALUES ('bool','Boolean value (yes / no)','');\nINSERT INTO `crm_field` VALUES ('email','E-mail input','');\nINSERT INTO `crm_field` VALUES ('enum','Single option picker','');\nINSERT INTO `crm_field` VALUES ('hidden','Hidden value','');\nINSERT INTO `crm_field` VALUES ('stamp','Date/time input','');\nINSERT INTO `crm_field` VALUES ('text','Text input','');\nINSERT INTO `crm_field` VALUES ('textarea','Text input (multi-line)','');\nPK\x07\x08f\x18\x1e\x84\xc5\x01\x00\x00\xc5\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x01\xa9HM\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00migrations.sqlUT\x05\x00\x012\xc7\xbb[CREATE TABLE IF NOT EXISTS `migrations` (\n `project` varchar(16) NOT NULL COMMENT 'sam, crm, ...',\n `filename` varchar(255) NOT NULL COMMENT 'yyyymmddHHMMSS.sql',\n `statement_index` int(11) NOT NULL COMMENT 'Statement number from SQL file',\n `status` varchar(16) NOT NULL COMMENT 'ok or full error message',\n PRIMARY KEY (`project`,`filename`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nPK\x07\x08%l\xb2\x16\x7f\x01\x00\x00\x7f\x01\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x01\xa9HME\x9b\x99\xdc{\x07\x00\x00{\x07\x00\x00\x1a\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x0020180704080000.base.up.sqlUT\x05\x00\x013\xc7\xbb[PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x01\xa9HMf\x18\x1e\x84\xc5\x01\x00\x00\xc5\x01\x00\x00%\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xcc\x07\x00\x0020180704080001.crm_fields-data.up.sqlUT\x05\x00\x013\xc7\xbb[PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x01\xa9HM%l\xb2\x16\x7f\x01\x00\x00\x7f\x01\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xed	\x00\x00migrations.sqlUT\x05\x00\x012\xc7\xbb[PK\x05\x06\x00\x00\x00\x00\x03\x00\x03\x00\xf2\x00\x00\x00\xb1\x0b\x00\x00\x00\x00"
	fs.Register(data)
}