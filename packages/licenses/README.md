## Package: Licenses

This packages allows to fetch SPDX and Choosealicense.com license information. This data is then stored as JSON files and exposed through typescript.
It is then used by Licensephobia to gather license metadata and also for the core logic of narrowing down
license rules.

The fetch tasks can be run periodically by workers to ensure data is up-to-date.