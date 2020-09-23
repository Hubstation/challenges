![](https://raw.githubusercontent.com/Hubstation/challenges/master/images/logo.png)

### Challenges are a great way to implement the skills learned via various platforms. This repository is focussed on creating an extensive set of challenges on various technologies by the community for the community.

#### Each Challenge will be part of a technology and the community can participate in two ways: 
     1 -  Creating Challenges
     2 -  Submitting solutions 
     

## Creating Challenges
![](https://raw.githubusercontent.com/Hubstation/challenges/master/images/challenge-diagram.png)

- Consider the above example for creating a challenge. A person who submits the challenge has to submit the solution as well.
- A Challenge will be part of a technology which will be a directory in the repo. In above case Kubernetes in the technology so any challanges related to Kubernetes should go under the Kubernetes directory.
- in above scenario C represents a challenge and S represents a solution, with "_B" or "_b" representing the difficulty level which is beginner in this case. Other levels include -> Intermediate - "_I" or "_i" and Expert - "_E" or "_e"

### Challenge format
To Keep things simple content inside the challenges can be in the form of a scenario with expectation. Like in above case it is a simple challenge to create a k3s cluster on raspberrypi. So you can have below portions in Challenge :

```
Challenge:  --------------
Expection: --------------
Architecture Diagram: if necessary
Other Details: if necessary
```

## Submitting Solutions
![](https://raw.githubusercontent.com/Hubstation/challenges/master/images/solution-diagram.png)

Apart from the "Challenge + solution" submitted by the person, there is an opportunity for the community to create their version of a solution and submit it as a Pull Request. This can happen inside the solutions folder where you will create a directory with your name and then post the solution with the same filename which in example above is "S_create_k3s_cluster_on_raspberrypi_solution". This way you can solve all challenges of a particular technology and have the solutions under your name.

Notes 
- To keep things Simple, Solutions can be mentioning a link to video, blog, git repo or just everything inside the solutions file. 
- Always create an Issue before submitting a PR



