ALTER TABLE users
ADD COLUMN id_role uuid references roles (id) on delete set default default uuid_nil();