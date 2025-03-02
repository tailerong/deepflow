/*
 * Copyright (c) 2024 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schema

const (
	DB_VERSION_TABLE    = "db_version"
	DB_VERSION_EXPECTED = "6.6.1.16"
)

const (
	CREATE_TABLE_DB_VERSION = `CREATE TABLE IF NOT EXISTS db_version (
		version             CHAR(64) PRIMARY KEY,
		created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at          DATETIME NOT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)ENGINE=innodb DEFAULT CHARSET=utf8;`
)
