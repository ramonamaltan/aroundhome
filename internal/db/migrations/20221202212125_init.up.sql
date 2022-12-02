create table if not exists partners(
    id              bigserial,
    servicename     varchar(255)    not null,
    latitude        float             not null,
    longitude       float             not null,
    radius          int             not null,
    constraint partners_pk
    unique (id)
);