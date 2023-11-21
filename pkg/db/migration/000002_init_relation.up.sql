ALTER TABLE "accounts" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");
ALTER TABLE "payments" ADD FOREIGN KEY ("owner") REFERENCES "users" ("email");
ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "products" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id");
ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "receipts" ADD FOREIGN KEY ("delivery_id") REFERENCES "deliveries" ("id");
ALTER TABLE "receipts" ADD FOREIGN KEY ("id") REFERENCES "orders" ("id");
ALTER TABLE "receipts" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");
ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_address" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_address" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");
ALTER TABLE "supplier_address" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");
ALTER TABLE "supplier_address" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id");
ALTER TABLE "product_in_cart" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");
ALTER TABLE "product_in_cart" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "product_in_order" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
ALTER TABLE "product_in_order" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "favorite_product" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "favorite_product" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");