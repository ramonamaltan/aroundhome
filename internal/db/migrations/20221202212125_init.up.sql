create table if not exists partners(
    id              bigserial       not null,
    partnername     varchar(255)    not null,
    servicename     varchar(255)    not null,
    latitude        float           not null,
    longitude       float           not null,
    material        varchar(255)    not null,
    radius          int             not null,
    rating          float           not null,
    constraint partners_pk
    unique (id)
);