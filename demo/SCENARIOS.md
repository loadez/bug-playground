# sem-agent Demo Scenarios

Each scenario is a self-contained prompt you paste into a clean Claude Code session.
The agent starts with zero context — only sem-agent installed and connected.

Project: `bug-playground` on Semaphore (loadez/bug-playground)

---

## Scenario 1: "Why did CI fail?"

**What it shows:** Agent cold-starts, finds the failure, reads logs, parses test results, identifies root cause — all via sem-agent. No browser needed.

**sem CLI comparison:** Would need `sem get workflows`, then `sem get pipelines`, then `sem logs` — and still get raw text, not parsed test results.

```
CI failed on the bug-playground project, main branch. Find out why and explain the root cause.
```

---

## Scenario 2: Self-Healing Loop

**What it shows:** Agent diagnoses, fixes the code, pushes, reruns, watches, and verifies — full autonomous loop.

**sem CLI comparison:** Impossible. sem CLI has no diagnose, no test parsing, no watch-and-verify.

```
CI is failing on bug-playground main branch. Diagnose the failures, fix the bugs in the code,
push the fix to a branch called "fix/discount-and-status", trigger a new workflow, and watch
it until it passes. Use sem-agent for all CI operations.
```

---

## Scenario 3: Pipeline Topology Analysis

**What it shows:** Dependency graph, critical path, blast radius — understand a pipeline structure without opening the UI.

**sem CLI comparison:** Not possible at all. sem CLI has no topology commands.

```
Analyze the pipeline topology for the latest workflow on bug-playground main branch.
Show me the critical path, and if there are failures, show the blast radius —
which blocks failed because of upstream failures vs which are root causes.
```

---

## Scenario 4: Flaky Test Detection

**What it shows:** Agent runs flaky detection across multiple builds, identifies non-deterministic tests.

**sem CLI comparison:** No test intelligence in sem CLI whatsoever.

```
Check if bug-playground has any flaky tests. Look at the last 10 workflow runs
and identify tests that sometimes pass and sometimes fail. Explain what makes them flaky.
```

---

## Scenario 5: Project Health Report

**What it shows:** Agent generates a health dashboard — pass rates, trends, deploy targets, overall verdict.

**sem CLI comparison:** Would require manually listing workflows and counting. No aggregation.

```
Generate a health report for the bug-playground project.
Include pass rates, recent failures, deployment target status, and an overall verdict.
```

---

## Scenario 6: Self-Discovery (Zero Docs)

**What it shows:** Agent discovers sem-agent capabilities on its own, no prior knowledge needed.

**sem CLI comparison:** sem has `--help` but no discovery, no examples, no capability map.

```
I have a CLI tool called sem-agent installed. I don't know what it does.
Explore its capabilities and tell me what I can do with it for CI/CD management.
Then check the status of a project called bug-playground.
```

---

## Scenario 7: Testbox — Local CI Environment

**What it shows:** Agent spins up a real Semaphore machine, syncs code, runs tests in actual CI env.

**sem CLI comparison:** `sem debug` exists but has no file sync, no composability, no agent integration.

```
I want to test my code in the real CI environment before pushing.
Use sem-agent testbox to warm up a machine for bug-playground,
sync the local code, and run the tests. Show me the results.
```

---

## Scenario 8: Multi-Branch Comparison

**What it shows:** Agent compares CI status across branches — useful for PR reviews.

**sem CLI comparison:** Manual per-branch lookup, no comparison.

```
Compare CI status between main and fix/discount-and-status branches on bug-playground.
Which branch is passing? What's different between them?
```

---

## Scenario 9: Deploy with Safety

**What it shows:** Dry-run promotion, then confirmed deploy, then wait for completion.

**sem CLI comparison:** `sem promote` is fire-and-forget. No dry run, no waiting.

```
The fix/discount-and-status branch is green on bug-playground.
First do a dry-run of promoting to "Deploy to Staging" (no --confirm).
Show me what would happen. Then if it looks good, actually promote and wait for it to complete.
```

---

## Scenario 10: Full Incident Response

**What it shows:** Complete triage → fix → deploy loop. The ultimate agent workflow.

**sem CLI comparison:** Each step manual, non-composable, requires browser for most of it.

```
Production alert: bug-playground CI is red on main. I need you to:
1. Diagnose what's wrong
2. Check if any failures are flaky (not real bugs)
3. For real bugs, fix the code
4. Push to a fix branch
5. Wait for CI to pass
6. Promote to staging
7. Give me a summary of everything you did
```

---

## Running Order (Recommended)

For a live demo, run in this order to build a narrative:

1. **Scenario 6** — Discovery (show sem-agent is self-documenting)
2. **Scenario 1** — Diagnose (show the problem)
3. **Scenario 3** — Topology (understand pipeline structure)
4. **Scenario 2** — Self-heal (fix it autonomously)
5. **Scenario 4** — Flaky detection (separate real bugs from noise)
6. **Scenario 5** — Health report (project overview)
7. **Scenario 9** — Safe deploy (promote with guardrails)

Scenarios 7, 8, 10 are bonus — use if time allows.

## Setup

```bash
# Install sem-agent
cd /path/to/sem-agent && make install

# Connect to Semaphore
sem-agent connect <org>.semaphoreci.com <token>

# Verify
sem-agent context show
```
