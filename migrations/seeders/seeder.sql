-- +migrage Up
INSERT INTO
    dd.users (
        name,
        email,
        password,
        phone,
        role_id,
        gender,
        birthdate,
        status
    )
VALUES
    (
        'Fajar',
        'fssalviro@gmail.com',
        'test',
        '1234',
        1,
        'male',
        '2021-01-23',
        'a'
    );

INSERT INTO
    dd.user_vendors (user_id, subdistrict_id, address, postal_code)
VALUES
    (1, 102, 'adderss', '412');

INSERT INTO
    dd.brands (user_vendor_id, name)
VALUES
    (1, 'distrodakwah');

INSERT INTO
    dd.departments (name)
VALUES
    ('departmenta');

INSERT INTO
    dd.subdepartments (department_id, name)
VALUES
    (1, 'subdepartmenta');

INSERT INTO
    dd.categories (subdepartment_id, name)
VALUES
    (1, 'category a');

INSERT INTO
    dd.product_kinds (name)
VALUES
    ('single product');

INSERT INTO
    dd.product_kinds (name)
VALUES
    ('variant produt');

INSERT INTO
    dd.product_types (value, name)
VALUES
    (1, 'consignment');

INSERT INTO
    dd.product_types (value, name)
VALUES
    (2, 'vendor');

INSERT INTO
    dd.product_types (value, name)
VALUES
    (3, 'consignment,vendor');

-- +migrate Down