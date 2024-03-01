# Overview

This **educational** project showcases the systematic exploration of pain-points within a fictional wine-making supply chain. Through a detailed examination, it unfolds the steps taken to identify inefficiencies and challenges, culminating in the determination that a blockchain-based system is the most suitable solution. The project will also showcase through working examples potential ways to implement solutions.

The project is presented from the viewpoint of a fictional company, codename `mengawas`, with the sole mission of assisting businesses in forming networks through the utilization of blockchain technologies. In instances where pronouns such as `we,` `our,` and `us` are used, they specifically refer to `mengawas`. View `mengawas` as a consultancy, presenting analysis which you will find in `docs` section and technology provider in the form of `working examples`. `mengawas` is not a participant in the supply chain itself.

For clarity and focus on problem-solving steps within the project, fictional entities, including countries (`Atlantis,` `Wakanda,` etc.), companies, and personas, are used in this project. This deliberate use of fictional elements and simplification of case studies aims to prevent distractions that might arise from the association with real-world entities. The intention is to inspire anyone who is interested in the project to find their own solution by learning from these methodologies and working examples. The working examples that we present here are purely illustrative and should not be copied entirely for used in a mission critical environment. 

## Context

The basic steps involved in wine making process is summarised in Figure 1.

![wine making](./assets/img/winemaking.webp)</br>
**Figure 1: Basic Wine Making Steps (see [Image source.](https://finding.wine/blogs/blog-posts/basic-steps-of-the-winemaking-process))**

The wine making process when viewed from a supply chain perspective is encapulated in these key stages: harvesting, processing, bottling, distribution, and selling (see Figure 2). We have bundled juice extraction, fermentation, pressure, filtration, and aging into a single stage known collectively as processing. The steps in the processing stage are seldom independently executed. Between the supply chain stages are transportations that serve as bridges between these stages, facilitating the movement of intermediate and final products.

Simultaneously, a regulatory chain runs parallel to the physical supply chain. Its primary role is to ensure compliance with regulations, particularly in areas such as health and safety, and product grading, at every stage of the physical chain.

![Supply chain](./assets/img/supplychain.jpg)</br>
**Figure 2: The supply chain**

### What problems are we trying to solve?

`Alantis` is a famous wine making country. The country's wine making supply chain is dominated by a single company, name `Alpha`, which control the processing, bottling and distribution of wines. `Alpha` also owns a vineyard and its chain of pubs. Other participants of wine making supply chain in `Atlantis` are dissatisfied with the situation.`Wakanda` a major importer of `Atlantis` wines is also dissatisfied with the dorminance of `Alpha`. Together with the supply chain participants in `Atlantis`, `Wakanda` lodge a complain to the `World Trading Body`.

The complains by `Wakanda` and other participants were accepted by the `World Trading Body`, which came to a judgement and ruled that `Alantis` must reform the wine supply chain make it more competitive. Allowing for more participants to compete in the supply chain. The ruling did not require the `Alpha` to be broken up. Only not to used its dominance to restrict others to compete.

`Atlantis` commissioned us to analyze the challenges to comply with the ruling of the `World Trade Body`. In response, we produced a report titled ["Tides of Complexity: Examining the Supply Chain Challenges in Atlantis"](./docs/challenge.md) for the `Atlantis` government. The report comprehensively details the current state of affairs.

### A proposed solution

Following from our analysis of the current state of affairs, we examined several options to address the challenges and came to the conclusion that a blockchain-based solution would be most appropriate. The methodologies we used to help reach our conclusion is detailed in the report titled ['Transforming Atlantis Wine Supply Chain: A Blockchain Solution.'](./docs/solution.md) we presented to the `Atlantis` government. In essence, the benefits of adopting blockchain technology in this context are:

* Facilitating peer-to-peer transactions among supply chain participants to make it easier for participants to share common data.
* Creating a network effect to enable economies of scale for participants while maintaining their independence to prevent anti-competitive practices.
* Providing regulatory bodies with enhanced oversight, enabling swift intervention in case of malpractices.

### Implementation roadmap

The implemention of our solution is discussed in our [roadmap](./docs/roadmap.md)

## Disclaimer

* The content of this project is intended for educational purpose only.
* The content is constantly updated and items may be removed and modified without warning.

## Copyright

Unless otherwise specified, the copyright in this project are assigned as follows.

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.