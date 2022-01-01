# GitDog 

[![Go](https://github.com/gitdog7/gitdog/actions/workflows/go.yml/badge.svg)](https://github.com/gitdog7/gitdog/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![CodeQL](https://github.com/gitdog7/gitdog/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/gitdog7/gitdog/actions/workflows/codeql-analysis.yml)

Get in-depth analysis reports of a GitHub repository.

# Examples

[rancher/rancher's Contributors Graph(force layout)](https://gitdog7.github.io/gitdog/examples/rancher/rancher/rancher_rancher_Contributors_Graph_force_.html)

[rancher/rancher's Contributors Graph(circular layout)](https://gitdog7.github.io/gitdog/examples/rancher/rancher/rancher_rancher_Contributors_Graph_circular_.html)

<table>
  <tr>
      <td width="50%" align="center">
      </td>
      <td width="50%" align="center">
      </td>
  </tr>
  <tr>
     <td><img src="docs/images/console.png"/></td>
     <td><img src="docs/images/project.png"/></td>
  </tr>
</table>

# Quick Start

## Build 
```shell
make build
```

## Run a Demo
```shell
bin/gitdog allinone -w workspace-go-echarts -r go-echarts/go-echarts
```

In this quick start demo, the **go-echarts/go-echarts** repository will be analyzed.

All the data will be stored in **workspace-go-echarts** folder.

A html file is generated in "./workspace-go-echarts/Contributors_and_Following_Relationships.html".
Open it in your Chrome, you will get this contributors' graph:

<img width="256" alt="截屏2021-12-26 下午9 42 23" src="https://user-images.githubusercontent.com/51254187/147410092-5e8e4ae7-bbf7-4304-94fc-6a30e923982e.png">

# A Step Further

## Generate a GitHub Token
It is strongly recommended that you obtain GitHub access token. 
If the GitHub token is not set, the rate limit allows for up to 60 requests per hour which is not enough for you to 
analysis a bit big github repository. 

https://github.com/settings/tokens/new

<img width="512" alt="截屏2021-12-25 下午11 06 45" src="https://user-images.githubusercontent.com/51254187/147388061-a04029f2-30a3-4374-af45-72cacf9ba6af.png">

Only "public_repo, read:org, read:user, repo:status" permission is required.

## Use allinone cmd to analysis a bigger repository (e.g. kubesphere/console)
```shell
bin/gitdog allinone -w workspace-ks -r kubesphere/console -t the_github_token_you_generated 
```

# Details 

# Contributions
