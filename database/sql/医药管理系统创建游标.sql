-- 查找所有的售出订单
CREATE PROCEDURE fetch_sale_order()
BEGIN
	DECLARE drug_sale_order_id INT;
	DECLARE employee_id INT;
	DECLARE sale_date DATETIME;
	DECLARE customer_id INT;
	DECLARE drug_name TEXT;
	DECLARE sale_price INT;
	DECLARE num INT;
	DECLARE total_price INT;
	
	DECLARE sale_order CURSOR
	FOR
	SELECT 
	dso.drug_sale_order_id , 
	dso.employee_id ,
	dso.sale_date ,
	dso.customer_id ,
	d.cname ,
	id.sale_price ,
	sdl.num ,
	id.sale_price * sdl.num as total_price
	FROM 
	drug_sale_order dso, 
	sale_drug_list sdl, 
	inventory_drug id, 
	drug d
	WHERE
	dso.drug_sale_order_id = sdl.drug_sale_order_id  AND 
	sdl.inventory_drug_id = id.inventory_drug_id AND 
	id.drug_id = d.drug_id ;
	
	OPEN sale_order;
	
	FETCH sale_order 
	INTO
	drug_sale_order_id, employee_id,sale_date,customer_id,drug_name,sale_price,num,total_price;
	SELECT 
	drug_sale_order_id, employee_id,sale_date,customer_id,drug_name,sale_price,num,total_price;
	
	CLOSE sale_order;
END//


-- 查找所有的进货订单
CREATE PROCEDURE fetch_buy_order()
BEGIN
	DECLARE drug_sale_order_id INT;
	DECLARE employee_id INT;
	DECLARE sale_date DATETIME;
	DECLARE customer_id INT;
	DECLARE drug_name TEXT;
	DECLARE sale_price INT;
	DECLARE num INT;
	DECLARE total_price INT;
	
	DECLARE sale_order CURSOR
	FOR
	SELECT 
	dso.drug_sale_order_id , 
	dso.employee_id ,
	dso.sale_date ,
	dso.customer_id ,
	d.cname ,
	id.sale_price ,
	sdl.num ,
	id.sale_price * sdl.num as total_price
	FROM 
	drug_sale_order dso, 
	sale_drug_list sdl, 
	inventory_drug id, 
	drug d
	WHERE
	dso.drug_sale_order_id = sdl.drug_sale_order_id  AND 
	sdl.inventory_drug_id = id.inventory_drug_id AND 
	id.drug_id = d.drug_id ;
	
	OPEN sale_order;
	
	FETCH sale_order 
	INTO
	drug_sale_order_id, employee_id,sale_date,customer_id,drug_name,sale_price,num,total_price;
	SELECT 
	drug_sale_order_id, employee_id,sale_date,customer_id,drug_name,sale_price,num,total_price;
	
	CLOSE sale_order;
END//


