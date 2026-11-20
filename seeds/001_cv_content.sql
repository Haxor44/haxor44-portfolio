-- =====================================================================
-- Portfolio content taken from the CV (ryan_matibo_malekya.pdf).
--
-- Run it (Postgres from docker-compose):
--   docker exec -i portfolio-postgres psql -U portfolio -d portfolio < seeds/001_cv_content.sql
--
-- Re-runnable: each block deletes the rows it owns (matched by title/name)
-- and inserts them fresh, all inside one transaction. Edit the values here
-- and re-run to update the site. Articles are NOT seeded — they come from
-- the Medium sync (MEDIUM_USERNAME in .env).
-- =====================================================================

BEGIN;

-- ---------------------------------------------------------------------
-- Projects
-- tech_stack is comma-separated; featured = true puts it on the home page.
-- ---------------------------------------------------------------------
DELETE FROM projects WHERE title IN (
  'DevSecOps CI/CD Pipeline & GitOps Deployment (E-MAT)',
  'AWS EKS Infrastructure Automation with Terraform',
  'On-Premise VM Automation with Proxmox & Terraform'
);

INSERT INTO projects
  (title, description, tech_stack, github_url, live_url, blog_url, deployment_platform, image_url, featured, display_order, created_at, updated_at)
VALUES
  ('DevSecOps CI/CD Pipeline & GitOps Deployment (E-MAT)',
   'Full DevSecOps pipeline for a 3-tier app: GitHub Actions CI with Trivy container scanning, and ArgoCD "App of Apps" GitOps across Dev, Staging and Prod with RBAC and namespace isolation. Monitored with Prometheus and Grafana, scaled with Redis caching and HPA.',
   'GitHub Actions,Trivy,ArgoCD,Kubernetes,Prometheus,Grafana,Redis',
   'https://github.com/Haxor44/cap-stone', NULL, NULL,
   'AWS EKS', NULL, true, 1, now(), now()),

  ('AWS EKS Infrastructure Automation with Terraform',
   'Production-grade AWS EKS cluster with Fargate and Managed Node Groups, provisioned entirely through Terraform. Least-privilege IAM policies and custom VPC networking secure cluster access, with lessons learned on dependency management and OIDC documented.',
   'Terraform,AWS EKS,Fargate,IAM,VPC',
   'https://github.com/Haxor44/terra-3', NULL,
   'https://medium.com/@evolmalek04/deploying-aws-eks-with-terraform-iam-policies-fargate-node-groups-and-lessons-learned-bf9ef0739906',
   'AWS EKS', NULL, true, 2, now(), now()),

  ('On-Premise VM Automation with Proxmox & Terraform',
   'Automated VM lifecycle management on Proxmox VE with Terraform and Cloud-Init, replacing manual provisioning with Infrastructure as Code. Reusable VM templates built with virt-customize standardise network and OS initialisation.',
   'Terraform,Proxmox VE,Cloud-Init,virt-customize',
   NULL, NULL,
   'https://medium.com/@evolmalek04/proxmox-cloud-init-setup-guide-automating-vm-deployment-with-terraform-40d1a1d052ba',
   'Proxmox VE', NULL, true, 3, now(), now());

-- ---------------------------------------------------------------------
-- Certifications — earned
-- The CV lists years only, so issue_date is 1 January of that year and
-- the site shows just the year. display_order keeps the CV's order
-- within the same year.
-- ---------------------------------------------------------------------
DELETE FROM certifications WHERE name IN (
  'AWS Certified AI Practitioner',
  'AI Fundamentals',
  'DevOps Engineering',
  'Kubernetes and Cloud Native Associate (KCNA)',
  'AWS Cloud Practitioner',
  'Cloud & Network Security',
  'DevOps Engineering Course',
  'Certified Kubernetes Application Developer (CKAD)',
  'HashiCorp Certified: Terraform Associate',
  'AWS Certified Solutions Architect – Associate'
);

INSERT INTO certifications
  (name, issuing_org, status, issue_date, display_order, created_at, updated_at)
VALUES
  ('AWS Certified AI Practitioner',               'Amazon Web Services',             'earned', '2026-01-01', 1, now(), now()),
  ('AI Fundamentals',                             'Zindua School',                   'earned', '2026-01-01', 2, now(), now()),
  ('DevOps Engineering',                          'Zindua School',                   'earned', '2025-01-01', 3, now(), now()),
  ('Kubernetes and Cloud Native Associate (KCNA)','CNCF / The Linux Foundation',     'earned', '2025-01-01', 4, now(), now()),
  ('AWS Cloud Practitioner',                      'Ajira AWS re/start Programme',    'earned', '2025-01-01', 5, now(), now()),
  ('Cloud & Network Security',                    'Cyber Shujaa',                    'earned', '2025-01-01', 6, now(), now()),
  ('DevOps Engineering Course',                   'Moringa School',                  'earned', '2024-01-01', 7, now(), now());

-- ---------------------------------------------------------------------
-- Certifications — in progress (no issue_date yet)
-- ---------------------------------------------------------------------
INSERT INTO certifications
  (name, issuing_org, status, issue_date, display_order, created_at, updated_at)
VALUES
  ('Certified Kubernetes Application Developer (CKAD)', 'CNCF / The Linux Foundation', 'in_progress', NULL, 1, now(), now()),
  ('HashiCorp Certified: Terraform Associate',          'HashiCorp',                   'in_progress', NULL, 2, now(), now()),
  ('AWS Certified Solutions Architect – Associate',     'Amazon Web Services',         'in_progress', NULL, 3, now(), now());

COMMIT;
