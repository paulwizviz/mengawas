# Requirements and Solution

This section describes the process `Mengawas` followed to define the problem, explore potential solutions, and translate the chosen architecture into actionable product backlogs.

## Problem Definition

Wine supply chain participants face several significant challenges:

* **Climate Change:** Increasingly unpredictable weather patterns lead to fluctuating yields and impact grape quality.
* **Information Silos:** The decentralized nature of the supply chain results in fragmented information sharing, hindering collaboration and traceability.
* **Varying Levels of Digitalization:** Participants have different levels of competence and investment in digital technologies, making a one-size-fits-all solution inappropriate.

## Solution Architecture Options

`Mengawas` aims to deliver a track and trace solution leveraging the Internet of Things (IoT) across the supply chain.  Two broad architectural options were considered:

* **Hub-and-Spoke Architecture:** (Figure 1)
* **Blockchain-Based Architecture:** (Figure 2)

<figure>
  <img src="../assets/img/hub-spoke.png" alt="Hub and spoke Solution" />
  <figcaption>Figure 1: Hub-and-Spoke Architecture</figcaption>
</figure>

**Hub-and-Spoke Architecture: Pros and Cons**

* **Pros:**
    * Can be managed by a single entity.
    * Follows a conventional three-tiered architecture (API, middleware, data).
* **Cons:**
    * Requires all participants to standardize on a single digital architecture.
    * Scaling can be challenging due to varying participant requirements.

<figure>
  <img src="../assets/img/blockchain.png" alt="Blockchain-based Solution" />
  <figcaption>Figure 2: Blockchain-Based Architecture</figcaption>
</figure>

**Blockchain-Based Architecture: Pros and Cons**

* **Pros:**
    * Well-suited for decentralized and independent participants.
    * Tokenization features offer participants options to monetize operations.
    * Supports peer-to-peer integration.
    * Customizable data security while facilitating data integration.
* **Cons:**
    * Does not follow a conventional three-tiered architecture.
    * May involve cryptocurrency, potentially complicating financing.

After analyzing the pros and cons of each option, `Mengawas` selected the blockchain-based solution for implementation.  This choice aligns best with the decentralized and diverse nature of the wine supply chain.

> **NOTE:**
> 1. A whiteboard architecture is an exercise in painting a broad picture of options available. It is not intended to focus on a specific solution.
> 2. Typically, several architectural options would be considered in a real-world scenario.
> 3. Architecture design is often an iterative process. The chosen architecture may evolve as circumstances change, even during implementation.

## Product Backlogs

The following product backlogs were identified based on the chosen blockchain architecture. Each backlog represents deliverables for a specific supply chain participant:

* **Victor Vineyard**
* **Vicky Vineyard**
* **Pete Processor**
* **Petra Processor**
* **Wendy Warehouse**
* **Mike Merchants**
* **Tommy Transporter**
* **George Grader**
* **Chris Custom**
* **Consumers**

A separate product backlog was also created for the blockchain infrastructure itself.

> **NOTE:**
> 1. In this context, a product backlog is a collection of epics, user stories, and chores representing the features and functionality required by a particular entity.
> 2. Given that each supply chain participant is independent and may have its own existing digital infrastructure, assigning individual product backlogs is appropriate.
> 3. The blockchain infrastructure, being an independent entity, requires its own dedicated product backlog.

[TO DO: Add links to resources on creating product backlogs and user stories]