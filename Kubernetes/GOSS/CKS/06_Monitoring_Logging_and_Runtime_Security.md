# Challenge: Develop a validating tooling to check against challenge solutions targted toward resolving the the CKS objective "Monitoring, Logging and Runtime Security"
## this could be 6 sub-tasks as described in the [CKS website](https://training.linuxfoundation.org/certification/certified-kubernetes-security-specialist/)
  * Perform behavioral analytics of syscall process and file activities at the host and container level to detect malicious activities
    - Test can check an outputfile for known artifacts that were setup and examinee was supposed to find them using any technique/tooling he knows
  * Detect threats within physical infrastructure, apps, networks, data, users and workloads
    - that is quite a load of testing, we can break it into more small tests for each, starting with the doable
  * Detect all phases of attack regardless where it occurs and how it spreads
    - Based on Mitre attack, create an attack scenarios and see if examinee found out which attack we used after saving it in an outfile or providing a mitigation
  * Perform deep analytical investigation and identification of bad actors within environment
    - againt this could be an output of sysdig inspect with known artifacts
  * Ensure immutability of containers at runtime
    - check if we can write
  * Use Audit Logs to monitor access
    - genertae audit events (simulation) and see responses.

Possible tooling: use [Goss](https://github.com/aelsabbahy/goss) for local tests, for network it might be extra challenging, however, lets start small.
