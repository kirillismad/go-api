create table
    "entities" (
        "id" integer primary key autoincrement,
        "uuid_field" text not null unique,
        "int_field" integer not null,
        "float_field" real not null,
        "datetime_field" datetime not null,
        "string_field" text not null,
        "bool_field" boolean not null,
        "json_field" json
    );