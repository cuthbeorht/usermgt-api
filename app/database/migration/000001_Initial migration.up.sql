-- Create user table
create table if not exists users (
    user_id uuid primary key not null,    
    email text,
    password text, 
    create_at timestamp,
    deleted_at timestamp
);