Here's a comprehensive `README.md` file for your PodLister project, complete with examples and output formatting. You can customize this template further to suit your needs.

# PodLister - Kubernetes Pod Management Tool

**PodLister** is a command-line tool written in Go that helps you manage Kubernetes pods in a specified namespace. It allows you to list pods, filter them based on their age, and perform actions such as exporting the list to a file or deleting old pods.

## Features

- **Namespace & Age Filtering:** List pods from any Kubernetes namespace and filter them based on their age.
- **Color-Coded Output:** Visualize pod information with color-coded age indicators for quick identification of old pods.
- **Pod Deletion:** Optionally delete pods that exceed a specified age threshold.
- **Formatted Output:** Display pod information in a neatly aligned table with columns for pod name, age, status, and node.
- **Export to File:** Save the list of pods to a CSV or JSON file for further analysis.

## Prerequisites

- Go (version 1.16 or later)
- Access to a Kubernetes cluster
- `kubectl` installed and configured to access your cluster

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/alianjo/podlister.git
   cd podlister
   ```

2. Build the application:

   ```bash
   go build -o podlister
   ```

3. (Optional) Move the binary to your PATH:

   ```bash
   mv podlister /usr/local/bin/
   ```

## Usage

To use PodLister, run the following command:

```bash
./podlister --namespace <namespace> --age <age> [--export <file>] [--delete]
```

### Command-Line Flags

- `--namespace`: Specify the Kubernetes namespace (default is `default`).
- `--age`: Specify the age threshold in days to filter old pods (default is `30`).
- `--export`: Export the result to a CSV or JSON file (e.g., `output.csv`).
- `--delete`: Delete pods older than the specified age threshold.

### Examples

#### Example 1: List Pods in the Default Namespace

```bash
./podlister --namespace default
```

**Output:**
```
NAME                                     AGE     STATUS    NODE
nfs-client-provisioner-75d5857d99-6tn4q  96 days Running   worker-2
nfs-client-provisioner-75d5857d99-c8dqh  50 days Running   worker-1
nginx                                    0 days  Pending   worker-1
```

#### Example 2: List Pods in a Specific Namespace

```bash
./podlister --namespace kube-system
```

**Output:**
```
NAME                                     AGE     STATUS    NODE
kube-dns-5648d7c6d9-f7hb8               50 days Running   master
coredns-78f78b9f8d-dnpl7                 30 days Running   master
```

#### Example 3: List Pods Older than 60 Days

```bash
./podlister --age 60
```

**Output:**
```
NAME                                     AGE     STATUS    NODE
nfs-client-provisioner-75d5857d99-6tn4q  96 days Running   worker-2
```

#### Example 4: Delete Pods Older than 30 Days

```bash
./podlister --age 30 --delete
```

**Output:**
```
Pod nfs-client-provisioner-75d5857d99-6tn4q deleted
```

#### Example 5: Export Pod List to a File

```bash
./podlister --export output.csv
```

**Output:**
```
Results exported to output.csv
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- [Kubernetes](https://kubernetes.io/)
- [Go](https://golang.org/)
- [Fatih Color](https://github.com/fatih/color)

### Summary of Sections:
- **Title and Description**: Brief overview of what the project is about.
- **Features**: List of key features of the tool.
- **Prerequisites**: Requirements for running the project.
- **Installation**: Instructions on how to clone, build, and optionally move the binary.
- **Usage**: How to use the command-line tool, including flags.
- **Examples**: Different usage examples with expected output.
- **Contributing**: Invitation for others to contribute.
- **License**: Licensing information.
- **Acknowledgments**: Recognition of libraries or resources used.

Feel free to modify this template based on your project specifics or preferences!