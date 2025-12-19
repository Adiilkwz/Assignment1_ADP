# Go Assignment 1: Management Systems Simulation

This repository contains a collection of four Go packages, each demonstrating core concepts of the Go programming language, including **Structs**, **Interfaces**, **Maps**, **Slices**, and **Method Receivers**.

## 🚀 Project Structure

The project is organized into modular packages to ensure clean code and separation of concerns:

```text
Assignment1/
├── main.go             # Entry point (Main Menu)
├── go.mod              # Module definition
├── Hotel/              # Task 1: Hotel Reservation System
├── Employee/           # Task 2: Employee & Interfaces
├── Gym/                # Task 3: Membership Management System
└── Wallet/             # Task 4: Digital Wallet Simulation

Features
Task 1: Hotel Room Reservation System
Concepts: Structs and Maps.

Functionality: Manage rooms using a map with the room number as the key. Includes methods for adding rooms, checking in/out, and filtering vacant rooms.

Task 2: Employee & Interfaces
Concepts: Interfaces, Polymorphism, and Slices.

Functionality: An Employee interface handles different employment types (FullTime, PartTime, Contractor, Intern).

Custom Logic: Each type implements its own CalculateSalary and a custom CalculateBonus logic.

Task 3: Membership Management System
Concepts: Interfaces and Maps.

Functionality: A gym system storing members in a map[uint64]Member. It demonstrates how to store different struct types (Basic and Premium) inside a single interface-based map.

Task 4: Wallet Transaction Simulation
Concepts: Pointer Receivers and Slices of Structs.

Functionality: Tracks a digital wallet balance and maintains a transaction history log. Methods include AddMoney and SpendMoney (with validation for insufficient funds).

How to Run
Ensure Go is installed:

go version

Navigate to the project root:

cd Assignment1

Run the program:

go run main.go

Implementation Details
Interfaces
Used in Task 2 and 3 to allow polymorphic behavior. This allows the program to call methods like GetDetails() or CalculateSalary() on a collection without knowing the specific underlying struct type.

Method Receivers
Value Receivers: Used for methods that only read data (e.g., GetBalance, GetDetails).

Pointer Receivers: Used for methods that modify the state of the object (e.g., AddMoney, CheckIn).