# Overview

While blockchain holds promise for industry transformation, it's vital to use it in ways that address, rather than worsen, industry challenges. In this project, we'll employ an illustrative approach to examine the challenges of a fictional wine-making supply chain and then systematically introduce blockchain to address those industry challenges.

We have deliberately chosen a fictional industry to prevent distractions that might arise from associations with real-world entities, such as copyright or intellectual rights violations. Nevertheless, the case study is rooted in elements of reality. The intention is to inspire you to draw from this project relevant techniques and working examples, and adapt them to your challenges.

## About this Project

In this project, we adopt the perspective of a fictional consultancy, code-named `mengawas`. When we use pronouns such as `we`, `our`, and `us`, they specifically refer to `mengawas`.

You'll also find a collection of demonstration blockchain and supporting technologies, such as Internet of Things (IoT), delivered from `mengawas` viewpoint.

## Project Context

The basic steps in the wine-making process are summarized in Figure 1.

![wine making](./assets/img/winemaking.webp)</br>
**Figure 1: Basic Wine Making Steps (see [Image source.](https://finding.wine/blogs/blog-posts/basic-steps-of-the-winemaking-process))**

When viewed from a supply chain perspective, the wine-making process is encapsulated in these key stages: harvesting, processing, bottling, distribution, and selling (see Figure 2). We have bundled juice extraction, fermentation, pressing, filtration, and aging into a single stage known collectively as `processing`. The steps in the processing stage are seldom independently executed. Between the supply chain stages are transportation segments that serve as bridges, facilitating the movement of intermediate and final products.

Simultaneously, a regulatory chain runs parallel to the physical supply chain. Its primary role is to ensure compliance with regulations, particularly in areas such as health and safety, and product grading, at every stage of the physical chain.

![Supply chain](./assets/img/supplychain.jpg)</br>
**Figure 2: The supply chain**

### What problems are we trying to solve?

`Atlantis` is a famous wine-producing country. The country's wine-making supply chain is dominated by a single company, named `Alpha`, which controls the processing, bottling, and distribution of wines. Alpha also owns a vineyard and a chain of pubs. Other participants in the wine-making supply chain in Atlantis are dissatisfied with the situation. `Wakanda`, a major importer of Atlantis wines, is also dissatisfied with the dominance of Alpha. Together with the supply chain participants in `Atlantis`, `Wakanda` lodges a complaint with the World Trading Body.

The complaints by `Wakanda` and other participants were accepted by the `World Trading Body`, which issued a judgment ruling that Atlantis must reform the wine supply chain to make it more competitive, allowing for more participants to compete in the supply chain. The ruling did not require Alpha to be broken up, only to refrain from using its dominance to restrict others from competing.

`Atlantis` commissioned us to analyze the challenges to comply with the ruling of the `World Trade Body`. In response, we produced a report titled ["Tides of Complexity: Examining the Supply Chain Challenges in Atlantis"](./docs/challenge.md) for the `Atlantis` government. The report comprehensively details the current state of affairs.

`Atlantis` commissioned us to analyze the challenges of complying with the ruling of the World Trade Body. In response, we produced a report titled ["Tides of Complexity: Examining the Supply Chain Challenges in Atlantis"](./docs/challenge.md) for the Atlantis government. The report comprehensively details the current state of affairs.

### A proposed solution

Following from our analysis of the current state of affairs, we examined several options to address the challenges and came to the conclusion that a blockchain-based solution would be most appropriate. The methodologies we used to help reach our conclusion is detailed in the report titled ['Transforming Atlantis Wine Supply Chain: A Blockchain Solution.'](./docs/solution.md) we presented to the `Atlantis` government. In essence, the benefits of adopting blockchain technology in this context are:

Following our analysis of the current state of affairs, we examined several options to address the challenges and concluded that a blockchain-based solution would be most appropriate. The methodologies we used to reach this conclusion are detailed in the report titled ['Transforming Atlantis Wine Supply Chain: A Blockchain Solution.'](./docs/solution.md), which we presented to the `Atlantis` government. In essence, the benefits of adopting blockchain technology in this context are:

* Facilitating peer-to-peer transactions among supply chain participants to make it easier for them to share common data.
* Creating a network effect to enable economies of scale for participants while maintaining their independence to prevent anti-competitive practices.
* Providing regulatory bodies with enhanced oversight, enabling swift intervention in case of malpractices.

### Implementation roadmap

The implemention of our solution is discussed in our [roadmap](./docs/roadmap.md)

## Disclaimer

* The content of this project is intended for educational purposes only.
* The entities mentioned in this project are purely fictional, and any resemblance to real entities is purely coincidental.
* The content is constantly updated, and items may be removed and modified without warning.

## Copyright

Unless otherwise specified, the copyright in this project are assigned as follows.

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.