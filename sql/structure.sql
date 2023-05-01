CREATE TABLE `car_characteristic`
(
    `id_car_characteristic` int(8) NOT NULL COMMENT 'id',
    `name`                  varchar(255)     DEFAULT NULL,
    `id_parent`             int(8)           DEFAULT NULL,
    `date_create`           int(10) UNSIGNED DEFAULT NULL,
    `date_update`           int(10) UNSIGNED DEFAULT NULL,
    `id_car_type`           int(8) NOT NULL,
    `name_eng`              varchar(255)     DEFAULT NULL COMMENT 'Название ENG',
    `name_pol`              varchar(255)     DEFAULT NULL COMMENT 'Название POL',
    `name_deu`              varchar(255)     DEFAULT NULL COMMENT 'Название DEU',
    `name_esp`              varchar(255)     DEFAULT NULL COMMENT 'Название ESP'
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Характеристики автомобилей';

CREATE TABLE `car_characteristic_value`
(
    `id_car_characteristic_value` int(8) NOT NULL,
    `value`                       varchar(255)     DEFAULT NULL,
    `unit`                        varchar(255)     DEFAULT NULL COMMENT 'Еденица измерения',
    `id_car_characteristic`       int(8) NOT NULL,
    `id_car_modification`         int(8) NOT NULL,
    `date_create`                 int(10) UNSIGNED DEFAULT NULL,
    `date_update`                 int(10) UNSIGNED DEFAULT NULL,
    `id_car_type`                 int(8) NOT NULL,
    `value_en`                    varchar(255)     DEFAULT NULL COMMENT 'Значение EN',
    `unit_en`                     varchar(255)     DEFAULT NULL COMMENT 'Единица измерения EN'
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Значения характеристик автомобиля';

CREATE TABLE `car_equipment`
(
    `id_car_equipment`    int(8)           NOT NULL COMMENT 'id',
    `name`                varchar(255)     NOT NULL,
    `date_create`         int(10) UNSIGNED NOT NULL,
    `date_update`         int(10) UNSIGNED NOT NULL,
    `id_car_modification` int(8)           NOT NULL,
    `price_min`           int(8) DEFAULT NULL COMMENT 'Цена от',
    `id_car_type`         int(8)           NOT NULL,
    `year`                int(8) DEFAULT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Комплектации';

CREATE TABLE `car_generation`
(
    `id_car_generation` int(8)           NOT NULL,
    `name`              varchar(255)     NOT NULL,
    `id_car_model`      int(8)           NOT NULL,
    `year_begin`        varchar(255)              DEFAULT NULL,
    `year_end`          varchar(255)              DEFAULT NULL,
    `date_create`       int(10) UNSIGNED NOT NULL,
    `date_update`       int(10) UNSIGNED          DEFAULT NULL,
    `id_car_type`       int(8)           NOT NULL DEFAULT '0'
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Поколения Моделей';

CREATE TABLE `car_mark`
(
    `id_car_mark` int(8)       NOT NULL COMMENT 'ID',
    `name`        varchar(255) NOT NULL,
    `date_create` int(10) UNSIGNED DEFAULT NULL,
    `date_update` int(10) UNSIGNED DEFAULT NULL,
    `id_car_type` int(8)       NOT NULL,
    `name_rus`    varchar(255)     DEFAULT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Марки автомобилей';

CREATE TABLE `car_model`
(
    `id_car_model` int(11)      NOT NULL COMMENT 'ID',
    `id_car_mark`  int(11)      NOT NULL,
    `name`         varchar(255) NOT NULL,
    `date_create`  int(10) UNSIGNED DEFAULT NULL,
    `date_update`  int(10) UNSIGNED DEFAULT NULL,
    `id_car_type`  int(8)       NOT NULL,
    `name_rus`     varchar(255)     DEFAULT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Модели автомобилей';

CREATE TABLE `car_modification`
(
    `id_car_modification`   int(11)      NOT NULL COMMENT 'ID',
    `id_car_serie`          int(11)      NOT NULL,
    `id_car_model`          int(11)      NOT NULL,
    `name`                  varchar(255) NOT NULL,
    `date_create`           int(10) UNSIGNED DEFAULT NULL,
    `date_update`           int(10) UNSIGNED DEFAULT NULL,
    `start_production_year` int(8)           DEFAULT NULL,
    `end_production_year`   int(8)           DEFAULT NULL,
    `id_car_type`           int(8)       NOT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Модификации автомобилей';

CREATE TABLE `car_option`
(
    `id_car_option` int(8)           NOT NULL,
    `name`          varchar(255)     NOT NULL,
    `id_parent`     int(8) DEFAULT NULL,
    `date_create`   int(10) UNSIGNED NOT NULL,
    `date_update`   int(10) UNSIGNED NOT NULL,
    `id_car_type`   int(8)           NOT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Опции';

CREATE TABLE `car_option_value`
(
    `id_car_option_value` int(8)           NOT NULL,
    `is_base`             tinyint(1)       NOT NULL DEFAULT '1',
    `id_car_option`       int(8)           NOT NULL,
    `id_car_equipment`    int(8)           NOT NULL,
    `date_create`         int(10) UNSIGNED NOT NULL,
    `date_update`         int(10) UNSIGNED NOT NULL,
    `id_car_type`         int(8)           NOT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Значения опций'
  ROW_FORMAT = COMPACT;

CREATE TABLE `car_serie`
(
    `id_car_serie`      int(11)      NOT NULL COMMENT 'ID',
    `id_car_model`      int(8)       NOT NULL,
    `name`              varchar(255) NOT NULL,
    `date_create`       int(10) UNSIGNED DEFAULT NULL,
    `date_update`       int(10) UNSIGNED DEFAULT NULL,
    `id_car_generation` int(8)           DEFAULT NULL,
    `id_car_type`       int(8)       NOT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Cерии автомобилей';


CREATE TABLE `car_type`
(
    `id_car_type` int(8)       NOT NULL,
    `name`        varchar(255) NOT NULL
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='Автомобильный сайт';