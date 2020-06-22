-- 卖药品订单总价
CREATE VIEW sale_drug_total_price(drug_sale_order_id, total_price) AS

SELECT d.drug_sale_order_id, SUM(i.sale_price*s.num) 
FROM inventory_drug AS i, drug_sale_order AS d, sale_drug_list AS s
WHERE d.drug_sale_order_id = s.drug_sale_order_id AND i.inventory_drug_id = s.inventory_drug_id
GROUP BY d.drug_sale_order_id;

-- 买药品订单总价
CREATE VIEW buy_drug_total_price(drug_buy_order_id, total_price) AS

SELECT dbo.drug_buy_order_id, SUM(sd.sale_price*bdl.num) 
FROM supplier_drug AS sd, drug_buy_order AS dbo, buy_drug_list AS bdl
WHERE dbo.drug_buy_order_id = bdl.drug_buy_order_id AND sd.drug_id = bdl.drug_id
GROUP BY dbo.drug_buy_order_id;

-- 药店每个订单从谁那买了什么药几个
CREATE VIEW buy_order_with_supplier_drug(drug_buy_order_id, supplier_id, supplier_name, drug_id, drug_name, num) AS
SELECT dbo.drug_buy_order_id , dbo.supplier_id, s.name , bdl.drug_id, d.cname, bdl.num 
FROM buy_drug_list bdl, drug_buy_order dbo, supplier s, drug d
WHERE bdl.drug_buy_order_id = dbo.drug_buy_order_id AND
      dbo.supplier_id=s.supplier_id AND
	  bdl.drug_id=d.drug_id;

-- 清单详情　包含　店员真实名称，顾客名，药品名，数量,　单价, 总价
CREATE VIEW drug_sale_order_detail(order_id, employee_id, employee_name, cust_id, cust_name, drug_id, drug_name, price, num, total_price) AS
SELECT dso.drug_sale_order_id , e.employee_id , e.name ,c.customer_id , c.name , d.drug_id , d.cname , id.sale_price , sdl.num ,  id.sale_price * sdl.num 
FROM employee e , customer c ,drug d ,inventory_drug id , drug_sale_order dso ,sale_drug_list sdl 
WHERE dso.drug_sale_order_id = sdl.drug_sale_order_id AND 
	  dso.employee_id = e.employee_id AND
	  dso.customer_id = c.customer_id AND 
	  sdl.inventory_drug_id = id.inventory_drug_id AND 
	  id.drug_id = d.drug_id;

