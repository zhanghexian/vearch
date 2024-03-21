#
# Copyright 2019 The Vearch Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied. See the License for the specific language governing
# permissions and limitations under the License.

# -*- coding: UTF-8 -*-

import requests
import json
import pytest
import logging
from utils.vearch_utils import *
from utils.data_utils import *

logging.basicConfig()
logger = logging.getLogger(__name__)

__description__ = """ test case for module alias """


class TestAlias:
    def setup(self):
        self.logger = logger

    def test_create_db(self):
        response = create_db(router_url, db_name)
        logger.info(response)
        assert response["code"] == 200

    @pytest.mark.parametrize(
        ["space_name"],
        [["ts_space"], ["ts_space1"]],
    )
    def test_create_space(self, space_name):
        embedding_size = 128
        space_config = {
            "name": space_name,
            "partition_num": 1,
            "replica_num": 1,
            "fields": [
                {
                    "name": "field_string",
                    "type": "keyword"
                },
                {
                    "name": "field_int",
                    "type": "integer"
                },
                {
                    "name": "field_float",
                    "type": "float",
                    "index": {
                        "name": "field_float",
                        "type": "SCALAR",
                    },
                },
                {
                    "name": "field_string_array",
                    "type": "string",
                    "array": True,
                    "index": {
                        "name": "field_string_array",
                        "type": "SCALAR",
                    },
                },
                {
                    "name": "field_int_index",
                    "type": "integer",
                    "index": {
                        "name": "field_int_index",
                        "type": "SCALAR",
                    },
                },
                {
                    "name": "field_vector",
                    "type": "vector",
                    "dimension": embedding_size,
                    "index": {
                        "name": "gamma",
                        "type": "FLAT",
                        "params": {
                            "metric_type": "InnerProduct",
                            "ncentroids": 2048,
                            "nsubvector": 32,
                            "nlinks": 32,
                            "efConstruction": 40,
                            "nprobe":80,
                            "efSearch":64,
                            "training_threshold":70000
                        }
                    },
                },
                # {
                #     "name": "field_vector_normal",
                #     "type": "vector",
                #     "dimension": int(embedding_size * 2),
                #     "format": "normalization"
                # }
            ]
        }

        response = create_space(router_url, db_name, space_config)
        logger.info(response)
        assert response["code"] == 200


    def test_create_alias(self):
        response = create_alias(router_url, "alias_name", db_name, space_name)
        logger.info(response)
        assert response["code"] == 200


    def test_get_alias(self):
        response = get_alias(router_url, "alias_name")
        logger.info(response)
        assert response["code"] == 200


    def test_update_alias(self):
        response = update_alias(router_url, "alias_name", db_name, "ts_space1")
        logger.info(response)
        assert response["code"] == 200

    def test_drop_alias(self):
        response = drop_alias(router_url, "alias_name")
        logger.info(response)
        assert response["code"] == 200

    def test_alias_array(self):
        response = create_alias(router_url, "alias_name1", db_name, space_name)
        assert response["code"] == 200

        response = create_alias(router_url, "alias_name2", db_name, space_name)
        assert response["code"] == 200

        response = get_all_alias(router_url)
        logger.info(response)
        assert response["code"] == 200

        response = drop_alias(router_url, "alias_name1")
        assert response["code"] == 200

        response = drop_alias(router_url, "alias_name2")
        assert response["code"] == 200

        response = get_all_alias(router_url)
        logger.info(response)
        assert response["code"] == 200

    @pytest.mark.parametrize(
        ["wrong_index", "wrong_type"],
        [
            [0, "create db not exits"],
            [1, "create space not exits"],
            [2, "update db not exits"],
            [3, "update space not exits"],
            [4, "get alias not exits"],
            [5, "delete alias not exits"],
            [6, "create alias exits"],
            [7, "update alias not exits"],
        ],
    )
    def test_alias_badcase(self, wrong_index, wrong_type):
        db_param = db_name
        if wrong_index == 0:
            db_param = "wrong_db"
            response = create_alias(router_url, "alias_name", db_param, space_name)
            logger.info(response)
            assert response["code"] != 200

        if wrong_index == 1:
            space_param = "wrong_space"
            response = create_alias(router_url, "alias_name", db_name, space_param)
            logger.info(response)
            assert response["code"] != 200

        if wrong_index == 2:
            db_param = "wrong_db"
            response = update_alias(router_url, "alias_name", db_param, space_name)
            logger.info(response)
            assert response["code"] != 200

        if wrong_index == 3:
            space_param = "wrong_space"
            response = update_alias(router_url, "alias_name", db_name, space_param)
            logger.info(response)
            assert response["code"] != 200

        if wrong_index == 4:
            response = get_alias(router_url, "alias_not_exist")
            logger.info(response)
            assert response["code"] != 200

        if wrong_index == 5:
            response = drop_alias(router_url, "alias_not_exist")
            logger.info(response)
            assert response["code"] != 200

        if wrong_index == 6:
            response = create_alias(router_url, "alias_name", db_name, space_name)
            assert response["code"] == 200
            response = create_alias(router_url, "alias_name", db_name, space_name)
            logger.info(response)
            assert response["code"] != 200
            response = drop_alias(router_url, "alias_name")
            assert response["code"] == 200

        if wrong_index == 7:
            response = update_alias(router_url, "alias_not_exist", db_name, space_name)
            logger.info(response)
            assert response["code"] != 200


    def test_destroy_db_and_space(self):
        space_info = list_spaces(router_url, db_name)
        for space in space_info["data"]:
            logger.info(drop_space(router_url, db_name, space["space_name"]))
        drop_db(router_url, db_name)