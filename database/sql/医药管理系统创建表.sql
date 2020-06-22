-- 经理表
CREATE TABLE manager
 (
    manager_id INT(5) AUTO_INCREMENT primary key,
    username VARCHAR(20) not null,
    password VARCHAR(10) not null,
    UNIQUE(username)
 );

 -- 员工表
 CREATE TABLE employee
 (
     employee_id  INT(5) AUTO_INCREMENT primary key,
     name VARCHAR(20) not null,
     username VARCHAR(20) not null,
     password VARCHAR(20) not null,
     position VARCHAR(20) not null,
     UNIQUE(username)
 );

-- 供应商
CREATE TABLE supplier
(
    supplier_id INT(5) AUTO_INCREMENT primary key,
    name VARCHAR(20) not null,
    phone VARCHAR(11) not null,
    UNIQUE(phone)
);

-- 顾客
CREATE TABLE customer
(
    customer_id INT(5) AUTO_INCREMENT primary key,
    name VARCHAR(20) not null,
    phone VARCHAR(11) not null,
    username VARCHAR(11) DEFAULT NULL COMMENT '用户名',
    password VARCHAR(20) DEFAULT '123456' COMMENT '密码',
    UNIQUE(phone)
);

-- 药品
CREATE TABLE drug
(
    drug_id INT(10) primary key,
    cname VARCHAR(50) not null COMMENT '中文名称',
    ename VARCHAR(50) COMMENT '英文名称',
    introduction text COMMENT '药品介绍',
    component text COMMENT '成分',
    property text COMMENT '性状',
    indication text COMMENT '适应症',
    medic_format text COMMENT '规格',
    taboo text COMMENT '禁忌',
    ytime text COMMENT '有效期',
    mstandard text COMMENT '执行标准',
    dosage text COMMENT '用法用量',
    adverseReactions text COMMENT '不良反应',
    interactions text COMMENT '药品相互作用',
    notice text COMMENT '注意事项',
    drug_type VARCHAR(50) COMMENT '药品类型',
    drug_health_type VARCHAR(50) COMMENT '药品医保类型',
    drug_recipe_type VARCHAR(50) COMMENT '药品处方类型'
);

-- 库存药
CREATE TABLE inventory_drug
(
    inventory_drug_id INT(10) AUTO_INCREMENT primary key,
    drug_id INT(10),
    purchase_price INT(5) not null,
    sale_price INT(5) not null,
    supplier_id INT(5),
    inventory_num INT(10) not null,
    foreign key (drug_id) references drug(drug_id),
    foreign key (supplier_id) references supplier(supplier_id)
);

-- 卖药品订单
CREATE TABLE drug_sale_order
(
    drug_sale_order_id INT(10) AUTO_INCREMENT primary key,
    employee_id INT(5),
    sale_date datetime not null,
    customer_id INT(5),
    foreign key (employee_id) references employee(employee_id),
    foreign key (customer_id) references customer(customer_id)
);

-- 卖药品清单
CREATE TABLE sale_drug_list
(
    sale_drug_list_id INT(10) AUTO_INCREMENT primary key,
    inventory_drug_id INT(10),
    num INT(10) not null,
    drug_sale_order_id INT(10),
    foreign key (inventory_drug_id) references inventory_drug(inventory_drug_id),
    foreign key (drug_sale_order_id) references drug_sale_order(drug_sale_order_id)
);

-- 供应商 药品 联系
CREATE TABLE supplier_drug
(
    supplier_drug_id INT(10) AUTO_INCREMENT primary key,
    supplier_id INT(5),
    drug_id INT(10),
    sale_price INT(10) not null,
    foreign key (drug_id) references drug(drug_id),
    foreign key (supplier_id) references supplier(supplier_id)
);

-- 买药品订单
CREATE TABLE drug_buy_order
(
    drug_buy_order_id INT(10) AUTO_INCREMENT primary key,
    manager_id INT(5),
    buy_date datetime not null,
    supplier_id INT(5),
    foreign key (supplier_id) references supplier(supplier_id),
    foreign key (manager_id) references manager(manager_id)
);

-- 买药品清单
CREATE TABLE buy_drug_list
(
    buy_drug_list_id INT(10) AUTO_INCREMENT primary key,
    drug_id INT(10),
    num INT(10) not null,
    drug_buy_order_id INT(10),
    foreign key (drug_id) references drug(drug_id),
    foreign key (drug_buy_order_id) references drug_buy_order(drug_buy_order_id)
);
