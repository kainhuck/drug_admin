-- 查找某年某个月内的所有销售订单
SELECT dso.drug_sale_order_id , d.cname , e.name ,c.name ,id.sale_price ,sdl.num , id.sale_price * sdl.num as total_price ,dso.sale_date 
FROM drug_sale_order dso ,sale_drug_list sdl ,employee e ,customer c ,inventory_drug id ,drug d 
WHERE dso.employee_id = e.employee_id  AND 
	  dso.customer_id = c.customer_id  AND 
	  dso.drug_sale_order_id = sdl.drug_sale_order_id  AND 
	sdl.inventory_drug_id = id.inventory_drug_id  AND id.drug_id = d.drug_id AND
    DATE(dso.sale_date) BETWEEN '2020-01-01' AND '2020-01-31';

-- 查找某年某个月内的所有进货订单
SELECT dbo.drug_buy_order_id ,d.cname , s.name , m.username , sd.sale_price ,bdl.num , bdl.num * sd.sale_price as total_price,dbo.buy_date 
FROM drug_buy_order dbo ,buy_drug_list bdl ,manager m ,drug d ,supplier s ,supplier_drug sd 
WHERE 
	dbo.drug_buy_order_id = bdl.drug_buy_order_id AND
	dbo.manager_id = m.manager_id AND 
	dbo.supplier_id = s.supplier_id AND
	bdl.drug_id = d.drug_id AND 
	sd.drug_id = d.drug_id AND
	sd.supplier_id = s.supplier_id AND
	DATE(dbo.buy_date) BETWEEN '2020-01-01' AND '2020-01-31';