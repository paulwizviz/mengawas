# Overview

This project demonstrates the use of peer-to-peer technologies, such as Blockchain, to support and enhance the wine-making supply chain. A supply chain refers to a network of participants involved in the production and delivery of a product.

### What does the wine-making supply chain look like?

The wine-making process consists of several essential stages (see Figure 1).

![wine making](./assets/img/winemaking.webp)  
**Figure 1: Basic Wine-Making Steps ([Image source](https://finding.wine/blogs/blog-posts/basic-steps-of-the-winemaking-process))**

From a supply chain perspective, the process can be broken down into key stages: harvesting, processing, bottling, distribution, and selling (see Figure 2). For simplicity, juice extraction, fermentation, pressing, filtration, and aging are grouped into a single stage called `processing`, as these steps are typically interrelated. Transportation connects these stages, moving both intermediate and final products, forming the production chain.

Alongside the production chain is the regulatory chain, which ensures compliance with regulations and quality standards.

The third component is the support chain, responsible for providing equipment (e.g., barrels, bottles) and services that enable the production chain to operate smoothly.

![Supply chain](./assets/img/supplychain.png)  
**Figure 2: The Wine-Making Supply Chain**

## Project Scenario

## Project Scenario

For this project, we have identified the following participants and their respective platform requirements (click on the links):

- [Delta Distiller](./docs/delta.md): Handles juice extraction, fermentation, pressing, filtration, aging, and bottling.  
- [Gerrard Grader](./docs/gerrard.md): Independent organization responsible for wine grading.  
- [Mike Merchant](./docs/mike.md): Retailer of the final product.  
- [Sierra Sensors](./docs/sierra.md): Provides sensor and satellite imagery services.  
- [Victor Vineyard](./docs/victor.md): Grape supplier.  

Each of these participants operates as an independent node within the supply chain, with its own platform. They must share information without relying on intermediaries to coordinate actions. The network is designed to be extensible, allowing new participants to join with minimal configuration.
## Project Scenario

For this project, we have identified the following participants and their respective platform requirements (click on the links):

- [Delta Distiller](./docs/delta.md): Handles juice extraction, fermentation, pressing, filtration, aging, and bottling.  
- [Gerrard Grader](./docs/gerrard.md): Independent organization responsible for wine grading.  
- [Mike Merchant](./docs/mike.md): Retailer of the final product.  
- [Sierra Sensors](./docs/sierra.md): Provides sensor and satellite imagery services.  
- [Victor Vineyard](./docs/victor.md): Grape supplier.  

Each of these participants operates as an independent node within the supply chain, with its own platform. They must share information without relying on intermediaries to coordinate actions. The network is designed to be extensible, allowing new participants to join with minimal configuration.

## Disclaimer

- This project is for demonstration purposes only.  
- The entities mentioned are fictional, and any resemblance to real entities is purely coincidental.  
- Content is subject to updates, and items may be removed or modified without prior notice.

## Copyright

Unless otherwise specified, the copyright for this project is assigned as follows:

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at:

<http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is provided "AS IS", without warranties or conditions of any kind. See the License for specific language governing permissions and limitations under the License.
