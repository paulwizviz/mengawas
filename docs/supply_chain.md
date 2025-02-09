# Supply Chain Analysis

A supply chain analysis identifies the key participants involved in producing and distributing a product, along with their respective roles.  

## The `mengawas` Supply Chain Analysis

For `mengawas`, this analysis helps us understand the flow of materials and information in the wine-making process, from grape to consumer, enabling us to pinpoint areas where improvements in efficiency and transparency can be made.  The simplified supply chain modeled in `mengawas` (Figure 1) serves as a basis to illustrate supply chain analysis.

<figure>
  <img src="../assets/img/supplychain.png" alt="Wine making supply chain" />
  <figcaption>Figure 1: The Wine-Making Supply Chain, illustrating the interconnectedness of the product, support, and regulatory chains.</figcaption>
</figure>

> **NOTE:**
> 1. This is a simplified, simulated supply chain created for the purposes of `mengawas`.
> 2. A real-world chain could involve more participants and more complex relationships than are depicted here.
> 3. The names of the participants are fictional.

## Product Chain

The product chain consists of participants directly responsible for the production of wine – i.e., harvesting, processing, warehousing, and selling.  "Processing" in `mengawas` encompasses juice extraction, fermentation, pressure treatment (for sparkling wines), filtration, aging, and bottling. These activities are crucial for ensuring wine quality.

The principal participants in the product chain and their relationship with each other are summarized in Figure 2.

<figure>
  <img src="../assets/img/product-chain.png" alt="Wine making supply chain" />
  <figcaption>Figure 2: The participants and their relationship in the product chain.</figcaption>
</figure>

The participants identified in Figure 2 are:

* **Victor Vineyard:** Supplies two grades of grapes (30% to Pete Processor for red wine and 70% to Petra Processor for white wine).
* **Vicky Vineyard:** Supplies grapes exclusively to Petra Processor.
* **Pete Processor:** Processes red wine grapes.
* **Petra Processor:** Processes white wine grapes.
* **Wendy Warehouse:** Stores bottled wine.
* **Mike Merchants:** Wine merchant with exclusive distribution rights to consumers.
* **Alice and Bob:** Consumers.

## Support Chain

> **NOTE:**
> 1. Real-world supply chains often involve multiple support chain participants, such as transportation companies, bottle suppliers, and providers of fermentation equipment.
> 2. Multiple participants are often involved in the distribution of semi-finished and finished products.  The focus on a single participant here is for simplicity.

The support chain includes participants not directly involved in production. `Mengawas` focuses on **Tommy Transporter**, responsible for transporting products within the product chain.

## Regulatory Chain

> **NOTE:** Real-world regulatory oversight often involves multiple regulatory agencies.  For simplicity, `mengawas` includes only two.

The regulatory chain consists of independent participants overseeing product safety and quality.  `mengawas` includes:

* **George Grader:** An agency that performs wine tasting and issues quality and appellation of origin certifications.
* **Chris Custom:** An agency that verifies compliance with excise taxes and customs regulations.