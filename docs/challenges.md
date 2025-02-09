# Challenges, Product Lifecycle Analysis, and Supply Chain Analysis

The initial step in addressing supply chain challenges, as in any software project, is to thoroughly understand the problem. This phase involves identifying challenges, performing a product lifecycle analysis, and conducting a supply chain analysis.

## Identify Challenges

This phase of the development lifecycle involves gathering input from all supply chain stakeholders to identify the challenges they face. Start by compiling a broad collection of challenges and listing them. Make this list available to all stakeholders. It is worth noting that nothing is fixed, and changes are expected as the development lifecycle proceeds.

### The `mengawas` case study

For case study purposes, let's imagine `mengawas` has identified these challenges for a simulated wine-making supply chain:

* **Climate Change:** Increasingly unpredictable weather patterns lead to fluctuating yields and impact grape quality.
* **Information Silos:** The decentralized nature of the supply chain results in fragmented information sharing, hindering collaboration and traceability.
* **Varying Levels of Digitalization:** Participants have different levels of competence and investment in digital technologies, making a one-size-fits-all solution inappropriate.

## Product Lifecycle Analysis

A product lifecycle analysis helps supply chain stakeholders understand the key data points needed to ensure product integrity and build consumer confidence. By examining each stage of a product's creation and distribution, we can identify crucial information that demonstrates quality and authenticity to the consumer. This focused approach ensures we collect and present the most relevant data, as illustrated in the following `mengawas` case study.

### The `mengawas` case study

Let's consider the wine production process. The steps, summarized in Figure 1, highlight the critical importance of lifecycle integrity for final product quality.

<figure>
  <img src="../assets/img/winemaking.webp" alt="Wine making step" />
  <figcaption>Figure 1: Basic Wine-Making Steps. <em>Source:</em> https://finding.wine/blogs/blog-posts/basic-steps-of-the-winemaking-process</figcaption>
</figure>

The roles of each step are:

1. **Harvest Grapes:** Grapes are harvested at optimal ripeness.
2. **Extract Juice:** Grapes are crushed and pressed to extract juice.
3. **Fermentation:** Yeast is added to the juice to convert sugars to alcohol.
4. **Pressure (Optional - Sparkling Wine):** A second fermentation creates bubbles in sparkling wines.
5. **Filtration:** The wine is filtered to remove solids.
6. **Aging:** The wine is aged, often in oak barrels or stainless steel tanks.
7. **Bottling:** The wine is bottled and sealed.
8. **Consumption:** The wine is ready to be enjoyed.

The following data points, in the context of `mengawas`, are crucial for both quality control during production and providing consumers with the information they need to make informed choices:

#### 1. Harvest & Grape Quality

| Data Point | Primary Use (Production/Consumer) | Details/Explanation |
|---|---|---|
| **Grape Maturity Metrics** | Production | Brix, pH, Titratable Acidity (TA), Berry Size & Weight.  These metrics are essential for winemakers to determine optimal harvest time and predict the wine's potential. |
| **Weather Data** | Production | Rainfall, temperature, and sunlight hours during the growing season significantly influence grape development and quality. |
| **Pesticide/Herbicide Application Records** | Production/Consumer | Transparency regarding chemical applications is crucial for consumer safety and builds trust, especially for organic and biodynamic wines. |
| **Grape Variety/Blend** | Consumer | Directly impacts flavor profile. *Example: "100% Chardonnay," "Blend of Merlot and Cabernet Franc"* |
| **Vintage (Year)** | Consumer | Reflects growing season conditions and influences quality. *Example: "2020"* |
| **Vineyard Location/Appellation** | Consumer | Indicates terroir and quality potential. Protected designations of origin (PDOs) offer further assurance. *Example: "Pauillac," "Willamette Valley"* |
| **Yield** | Production | Tons/acre or kg/hectare. |
| **Organic/Biodynamic Certification** | Consumer | Verifies sustainable practices. *Example: "Certified Organic," "Biodynamic"* |

#### 2. Juice Extraction & Fermentation

| Data Point | Primary Use (Production/Consumer) | Details/Explanation |
|---|---|---|
| **Crushing/Pressing Method** | Production | Influences tannin and phenolic extraction. |
| **Must Analysis** | Production | Sugar, acidity, pH, yeast nutrients. |
| **Fermentation Temperature & Duration** | Production | Crucial for flavor and aroma development. |
| **Fermentation Vessel** | Production | Stainless steel, oak, etc. |
| **Malolactic Fermentation** | Production | If applicable, softens acidity. |
| **Sulfur Dioxide (SO2) Levels** | Production | Monitored for preservation. |
| **Alcohol Content** | Consumer | *Example: "14.5% ABV"* |
| **Yeast Strain(s) Used** | Production/Enthusiast | Impacts fermentation characteristics and flavor. *Example: "Using a specific strain of Saccharomyces cerevisiae"* |

#### 3. Aging & Maturation

| Data Point | Primary Use (Production/Consumer) | Details/Explanation |
|---|---|---|
| **Aging Vessel Type & Size** | Production | Oak barrels, stainless steel tanks, etc. |
| **Oak Type & Toast Level** | Production | Contributes specific flavor compounds. |
| **Aging Duration** | Production/Enthusiast | Influences complexity. *Example: "Aged for 18 months in French oak barrels"* |
| **Temperature & Humidity of Aging Cellar** | Production | Affects aging rate. |
| **Racking/Filtering/Fining Practices** | Production | Clarification and stabilization methods. |
| **Microbial Analysis** | Production | Checks for spoilage organisms. |

#### 4. Bottling & Packaging

| Data Point | Primary Use (Production/Consumer) | Details/Explanation |
|---|---|---|
| **Bottling Date** | Consumer | Useful for tracking and understanding potential aging. *Example: "Bottled on June 10, 2022"* |
| **Filtration Method** | Production | Influences clarity and potential flavor changes. |
| **Cork/Closure Type** | Production | Affects aging potential and oxygen exchange. *Example: "Natural cork," "Screw cap"* |
| **Bottle Type & Size** | Production | Impacts aging and presentation. |
| **Quality Control Checks** | Production | Fill level, cork integrity, visual inspection. |

#### 5. Distribution & Storage

| Data Point | Primary Use (Production/Consumer) | Details/Explanation |
|---|---|---|
| **Storage Temperature & Humidity** | Production | During transport and retail. |
| **Shipping Records/Provenance Information** | Consumer | Provides authenticity and traceability. *Example: "Shipped directly from the winery to the consumer"* |

#### 6. Final Product Information (Consumer)

| Data Point | Primary Use (Production/Consumer) | Details/Explanation |
|---|---|---|
| **Label Information** | Consumer | Accuracy and compliance are essential. |
| **Winemaker's Notes** | Consumer | Tasting notes, winemaking techniques, and the winemaker's philosophy add context and enhance consumer understanding. *Example: "This full-bodied Cabernet Sauvignon exhibits notes of black currant, cedar, and vanilla, with a long, smooth finish."* |
| **Certifications** | Consumer | Organic, biodynamic, sustainable, or other relevant certifications. *Example: "Certified Sustainable Wine"* |

> **NOTE:**
> 1. The wine-making process is simplified as `mengawas` is an educational project.
> 2. A comprehensive analysis for a real-world production scenario would require thorough interviews with multiple sources and consideration of different wine types.
> 3. The data points presented are illustrative and presented in tabular form, but for a real-world project, the analysis can be presented in other forms.
> 4. Ensure that the result of the analysis is shared with all relevant stakeholders.

## Supply Chain Analysis

It is rare for any product manufacturing process to be undertaken completely by a single entity. In many real-world production processes, there are multiple participants operating independently but collaborating in the form of a supply chain. Hence, the purpose of the supply chain analysis is to identify the key participants involved in producing and distributing a product, along with their respective roles. 

### The `mengawas` case study

The following `mengawas` case study, based on a simulated wine-making supply chain, demonstrates how this analysis can be applied in practice.

<figure>
  <img src="../assets/img/supplychain.png" alt="Wine making supply chain" />
  <figcaption>Figure 1: The Wine-Making Supply Chain, illustrating the interconnectedness of the product, support, and regulatory chains.</figcaption>
</figure>

The wine-making supply chain is composed of three sub-chains:

* [Product chain](#product-chain)
* [Support chain](#support-chain)
* [Regulatory chain](#regulatory-chain)

> **NOTE:**
> 1. The names of the participants are fictional.
> 2. The product, support, and regulatory chains are applicable to all supply chain scenarios.

### Product chain

The product chain consists of participants directly responsible for the production of wine – i.e., harvesting, processing, warehousing, and selling. "Processing" in `mengawas` encompasses juice extraction, fermentation, pressure treatment (for sparkling wines), filtration, aging, and bottling. In real-world scenarios, these processing steps might be divided among multiple participants. For this use case, the assumption is "processing" is done by one participant.

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

Having identified the key participants, each is further analyzed in detail. This analysis includes the following parameters:

* Identify personas (humans) in each participant and their responsibilities particularly in relation to data collections.
* Its relationship to other participants in the supply chain.
* Product data points that this participant is responsible for.
* Lists of challenges the participant faces.

The following is an example of the analysis of **Victor Vineyard**.

* **Victory Vineyard** 
  * Personas:
    * John the Farmer
    * Jane the data analyst
  * Relationship to other participants:
    * George Grader (issue operating license)
    * Tommy Transporter (responsible for transporting grapes)
  * Data points:
    * Soil pH value, 
    * Berry Size & Weight
  * Challenges:
    * John has to collect data manually
    * Jane lacks access to data in a format suitable for automated analysis.

A similar analysis would be conducted for other participants but it is not shown here for brevity.

> **NOTE:**
> 1. The suggested parameters are a starting point; a real-world analysis may require additional data points.
> 2. The information are collected as bullet points but it can be presented in other format.

### Support Chain

The support chain includes participants not directly involved in production. `Mengawas` focuses on **Tommy Transporter**, responsible for transporting products within the product chain.

> **NOTE:**
> 1. Real-world supply chains often involve multiple support chain participants, such as transportation companies, bottle suppliers, and providers of fermentation equipment.
> 2. Multiple participants are often involved in the distribution of semi-finished and finished products. The focus on a single participant here is for simplicity.

### Regulatory Chain

The regulatory chain consists of independent participants overseeing product safety and quality. `Mengawas` includes:

* **George Grader:** An agency that performs wine tasting and issues quality and appellation of origin certifications.
* **Chris Custom:** An agency that verifies compliance with excise taxes and customs regulations.

> **NOTE:** Real-world regulatory oversight often involves multiple regulatory agencies. For simplicity, `mengawas` includes only two.

## The Iterative Nature of the Analysis

The process of identifying challenges, analyzing the product lifecycle, and understanding the supply chain is rarely a linear, one-time activity. As the project progresses, new information may emerge, priorities may shift, and the understanding of the problem may evolve. Therefore, these analyses must be revisited and refined throughout the development lifecycle.

The use cases as depicted in this page is only to illustrate one way of capturing results of the analysis. 