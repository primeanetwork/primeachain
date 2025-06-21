# Server Prepping Role

## Overview

The `server-prepping` role is the foundational role in the Primea blockchain deployment automation. It prepares Ubuntu servers with the essential system configuration and dependencies required for running Primea blockchain nodes.

## Purpose in Primea Project

This role serves as the first step in the deployment pipeline, ensuring all servers have:
- Updated system packages for security and stability
- Essential development tools for building and running Go applications
- Proper firewall configuration for security
- Consistent timezone settings across all nodes
- Base directory structure for Primea blockchain deployment

## What This Role Does

### 1. System Updates
- Updates the apt package cache
- Upgrades all installed packages
- Removes unnecessary packages and cleans up

### 2. Package Installation
Installs the following essential packages:
- **git**: Version control for code management
- **make**: Build automation tool
- **golang-go**: Go programming language runtime
- **curl**: HTTP client for API interactions
- **ufw**: Uncomplicated Firewall for network security
- **htop**: System monitoring tool
- **build-essential**: Essential compilation tools

### 3. Security Configuration
- Enables UFW firewall with deny-by-default policy
- Allows only OpenSSH connections for secure remote access

### 4. System Configuration
- Sets timezone to UTC for consistency across all nodes
- Creates `/opt/primea` directory with proper ownership
- Ensures ubuntu user exists with proper shell

## Usage

This role should be applied to all servers in your inventory before deploying any blockchain-specific roles.

### Example Playbook Usage

```yaml
---
- hosts: all
  roles:
    - server-prepping
```

### Inventory Groups

This role is designed to work with the following inventory structure:
- `public-rpc`: Public RPC nodes
- `private-rpc`: Private RPC nodes  
- `validators`: Validator nodes

## Requirements

- Ubuntu 20.04+ servers
- Ansible 2.9+
- Sudo privileges for the ubuntu user
- Internet connectivity for package installation

## Dependencies

Currently, this role has no dependencies. It is designed to be the first role executed in the deployment pipeline.

## Variables

No custom variables are required for this role. All configuration uses sensible defaults.

## Handlers

This role does not define any handlers.

## Tags

No specific tags are defined for this role.

## Example Output

When successfully executed, this role will:
- Update and upgrade all system packages
- Install all required dependencies
- Configure UFW firewall with OpenSSH access only
- Set timezone to UTC
- Create `/opt/primea` directory owned by ubuntu:ubuntu

## Security Considerations

- UFW is configured with a deny-by-default policy
- Only OpenSSH is allowed through the firewall
- All packages are updated to latest versions for security patches
- The ubuntu user is created with proper permissions

## Troubleshooting

### Common Issues

1. **Package installation fails**: Ensure the server has internet connectivity
2. **UFW configuration fails**: Check if UFW is already configured
3. **Permission denied**: Ensure the ubuntu user has sudo privileges

### Debug Mode

To run this role in debug mode:
```bash
ansible-playbook -i inventory playbook.yml -vvv
```

## Contributing

When modifying this role:
1. Test changes on a development environment first
2. Ensure all tasks are idempotent
3. Update this README if new functionality is added
4. Follow Ansible best practices for role development 
