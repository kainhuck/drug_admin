-- 在新增用户订单前检查药品数量
CREATE TRIGGER b_i_sale_drug_list BEFORE INSERT ON sale_drug_list FOR EACH ROW 
BEGIN 
	DECLARE inv INT;
	SELECT id.inventory_num INTO @inv FROM inventory_drug id WHERE id.inventory_drug_id = NEW.inventory_drug_id;
	IF NEW.num > @inv 
    THEN INSERT INTO XXX VALUES('XXX');
	END IF;
END//

-- 在新增用户订单后减少药品数量
CREATE TRIGGER a_i_sale_drug_list AFTER INSERT ON sale_drug_list FOR EACH ROW 
BEGIN 
	UPDATE inventory_drug SET inventory_num = inventory_num - NEW.num WHERE inventory_drug_id = NEW.inventory_drug_id;
END//

-- 新增进货订单后增加药品数量
/* CREATE TRIGGER a_i_buy_drug_list AFTER INSERT ON buy_drug_list FOR EACH ROW 
BEGIN 
	UPDATE inventory_drug SET inventory_num = inventory_num + NEW.num WHERE drug_id = NEW.drug_id AND supplier_id=NEW.supplier_id;
END// */

