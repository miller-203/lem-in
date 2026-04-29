Here is a professional and comprehensive README.md file tailored for your lem-in project. It follows the standard documentation style expected in the 01Edu / Zone01 ecosystem.

Lem-in
Lem-in is a digital ant farm simulation developed in Go. The goal of the project is to find the most efficient way to move a specific number of ants across a network (colony) of rooms and tunnels.

🐜 The Challenge
The program reads a description of an ant farm from a file. This colony consists of rooms linked by tunnels. The objective is to navigate all ants from the ##start room to the ##end room in the fewest number of turns possible.

Key Constraints:
One Ant per Room: Only one ant can occupy a room at a time (except for ##start and ##end).

Optimal Pathfinding: The shortest path is not always the best. When dealing with many ants, the program must find multiple non-overlapping paths to avoid "traffic jams."

Turn-based Movement: Each ant can move only once per turn through a single tunnel.

🛠️ Features
Robust Parsing: Handles complex input files including room coordinates, links, and comments.

Algorithm Optimization: Implements advanced graph theory (e.g., Breadth-First Search or flow-based algorithms like Edmonds-Karp) to identify the optimal path combination.

Error Handling: Gracefully manages invalid data formats (invalid ant counts, missing start/end rooms, duplicated links, etc.) and outputs ERROR: invalid data format.

Standard Output Compliance: Displays the original map followed by the specific movements of each ant in the format Lx-y (where x is the ant ID and y is the destination room).
