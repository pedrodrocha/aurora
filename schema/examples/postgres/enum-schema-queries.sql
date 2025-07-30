-- How enums appear in PostgreSQL schema queries

-- 1. View all custom enum types
SELECT t.typname as enum_name,
       e.enumlabel as enum_value
FROM pg_type t 
JOIN pg_enum e ON t.oid = e.enumtypid  
WHERE t.typname = 'user_role_type'
ORDER BY e.enumsortorder;

-- 2. View table column information
SELECT 
    c.column_name,
    c.data_type,
    c.udt_name,
    c.is_nullable
FROM information_schema.columns c
WHERE c.table_name = 'user_roles'
  AND c.column_name = 'role';

-- 3. Alternative view using pg_catalog
SELECT 
    a.attname as column_name,
    t.typname as data_type,
    a.attnotnull as not_null
FROM pg_attribute a
JOIN pg_type t ON a.atttypid = t.oid
JOIN pg_class c ON a.attrelid = c.oid
WHERE c.relname = 'user_roles'
  AND a.attname = 'role'
  AND a.attnum > 0;

-- Expected results:
-- Query 1: Shows enum values: admin, user, moderator
-- Query 2: data_type = 'USER-DEFINED', udt_name = 'user_role_type'  
-- Query 3: data_type = 'user_role_type'