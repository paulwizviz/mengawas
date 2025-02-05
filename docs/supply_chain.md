# Supply Chain

`Mengawas` models the wine making supply chain, as illustrated in Figure 1.

<figure>
  <img src="./assets/img/supplychain.png" alt="Wine making supply chain" />
  <figcaption>Figure 1: The Wine-Making Supply Chain, illustrating the interconnectedness of the product, support, and regulatory chains.</figcaption>
</figure>

In the figure, `mengawas` has identified three sub-chains:

* [Product chain](#product-chain)
* [Support chain](#support-chain)
* [Regulatory chain](#regulatory-chain)

## Product Chain 

The product chain consists of participants directly responsible for the production of wine - i.e. `harvesting`, `processing`, `warehousing`, and `selling`. In the context of `mengawas`, `processing` refers to the following activities:

* Extract Juice
* Fermentation
* Pressure
* Filtration
* Aging
* Bottling

The principal participants in the product chain and their relation to each other is summarised in Figure 2.

<figure>
  <img src="./assets/img/product-chain.png" alt="Wine making supply chain" />
  <figcaption>Figure 2: The participants and their relationship in the product chain.</figcaption>
</figure>

> NOTE: 
>1. This is a simulated chain simplified for `mengawas`. A real world chain could involve more participants than is depicted.
>2. The names of the participants are fictional. 

As per Figure 2, the participants identified are:

* **Victor Vineyard** - supplies two grades of grapes (30% to Pete Processor and 70% to Petra Processor)
* **Vicky Vineyard** - only supply to Petra processor
* **Pete Processor** - Process only Red wine
* **Petra Processor** - Process only White wine
* **Wendy Warehouse** - Takes bottles from Pete and Petra for storage
* **Mike Merchants** - Has a monopoly to supply products to consumer
* **Alice and Bob** - Consumers

## Support Chain

The support chain consists of participants the are not directly involved in the production of wines -- e.g. transporting products between production participants.

## Regulatory Chain

The regulatory chain consists of participants that are independent of the product chain overseeing the safety and quality of the product.

