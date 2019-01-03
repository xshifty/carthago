# Carthago domain model definition

The basic concept of this software is to provide a simple way to sell your products through the Internet.
For the first version, I plan to implement support for **Basic Stock Management**, **Shopping Cart** and **Order Handling**.

## Basic Stock Management

To update our **Product** stocks, we have to create a **Acquired Batch** from a **Supplier**.
If the **Product** and the **Supplier** already exists in the system, we just have to create new **Acquired Batch** bound to existing **Supplier** and **Product**.

To verify the current available amount of a **Product** in stock, we have to build a **Stock Detail** based on previous **Acquired Batch** and **Selling Order**.
