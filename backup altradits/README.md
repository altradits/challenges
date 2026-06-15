<div align="center">

<img src="logo.png" width="100%" alt="Altradits"/>

# Stanley Chege Thuita

<br/>

<a href="https://github.com/altradits">
  <img src="https://readme-typing-svg.demolab.com/?font=Fira+Code&weight=500&size=22&duration=3200&pause=1000&color=FF6B6B&center=true&vCenter=true&width=720&height=50&lines=Bitcoin+Lightning+%2B+M-Pesa+wallet+backend+in+Go;12+months+deep+in+Go+%26+Lightning;SE+Apprentice+%40+Zone01+Kisumu;Open+to+Backend+%26+Bitcoin%2FLightning+roles" alt="Typing SVG"/>
</a>

<br/>

[![Profile Views](https://komarev.com/ghpvc/?username=altradits&style=for-the-badge&color=274C77&label=PROFILE+VIEWS)](https://github.com/altradits)
[![Open to Work](https://img.shields.io/badge/STATUS-OPEN_TO_WORK-2EA043?style=for-the-badge&logo=target&logoColor=white)](mailto:altradits@gmail.com)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Stanley_Chege_Thuita-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/stanmobitech)
[![Email](https://img.shields.io/badge/Email-altradits%40gmail.com-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:altradits@gmail.com)
[![Location](https://img.shields.io/badge/Based_in-Kisumu%2C_Kenya-274C77?style=for-the-badge&logo=googlemaps&logoColor=white)](https://www.zone01kisumu.ke/)

</div>

<img src="https://capsule-render.vercel.app/api?type=rect&color=274C77&height=3&width=100&section=footer" width="100%"/>

## ⚡ Snapshot

From "hello world" to a working wallet backend, one shipped phase at a time.

- 🎓 SE Apprentice @ **[Zone01 Kisumu](https://www.zone01kisumu.ke/)**, peer-to-peer and project-based, no lectures
- ⚡ Founder & sole developer of **[Altradits](#-altradits--bitcoin-banking-for-kenya)** — Bitcoin banking for Kenya, M-Pesa rails, vanilla Go
- 🧩 Owns features end to end, from schema to deploy

<img src="https://github-readme-stats.vercel.app/api?username=altradits&show_icons=true&theme=tokyonight&hide_border=true&bg_color=0D1117&title_color=274C77&icon_color=FF6B6B&text_color=c9d1d9&count_private=true&include_all_commits=true" width="100%" alt="stats"/>

<img src="https://capsule-render.vercel.app/api?type=rect&color=FF6B6B&height=3&width=100&section=footer" width="100%"/>

## 🧰 Toolbox

**Core**, vanilla first while I build solid fundamentals:

<img src="https://skillicons.dev/icons?i=go,postgres,redis,git,github,linux,bash,docker&theme=dark" alt="core skills"/>

**Exploring** in side projects:

<img src="https://skillicons.dev/icons?i=nextjs,ts,tailwind,nodejs&theme=dark" alt="exploring"/>

<img src="https://github-readme-stats.vercel.app/api/top-langs/?username=altradits&layout=compact&theme=tokyonight&hide_border=true&bg_color=0D1117&title_color=274C77&text_color=c9d1d9&langs_count=8" width="100%" alt="top languages"/>

<img src="https://capsule-render.vercel.app/api?type=rect&color=274C77&height=3&width=100&section=footer" width="100%"/>

## 🏗️ Altradits — Bitcoin Banking for Kenya

> "The root problem with conventional currency is all the trust that's required to make it work." — Satoshi Nakamoto, 2008

Everyone in my community has a story — an "nguru" who vanished with their savings, a Sacco that quietly went broke, a chama that fell apart when the treasurer disappeared. Altradits is my answer: a bank for my community where there's no one left to trust — only code, Bitcoin, and a ledger anyone can audit. Deposit the M-Pesa you already use, lock your sats for 5+ years, and let patience do what no nguru ever could.

[![Repo](https://img.shields.io/badge/Repo-altradits%2Faltradits-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/altradits/altradits)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-2EA043?style=for-the-badge)](https://github.com/altradits/altradits/blob/main/LICENSE)

<p>
<img src="https://img.shields.io/badge/Bitcoin-000000?style=for-the-badge&logo=bitcoin&logoColor=F7931A" alt="Bitcoin"/>
<img src="https://img.shields.io/badge/Lightning_Network-792EE5?style=for-the-badge&logo=lightning&logoColor=white" alt="Lightning Network"/>
<img src="https://img.shields.io/badge/LND-2A2A2A?style=for-the-badge&logo=bitcoinsv&logoColor=F7931A" alt="LND"/>
<img src="https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white" alt="JWT"/>
<img src="https://img.shields.io/badge/M--Pesa-34A853?style=for-the-badge&logo=mpesa&logoColor=white" alt="M-Pesa"/>
</p>

- ⚡ Send/receive sats over Lightning + `you@altradits.com` addresses (LNURL-pay)
- 🇰🇪 Deposit & withdraw KES via M-Pesa STK push
- 💰 Treasury pool: bonds, money market, equities, BTC, auto interest
- 🛡️ JWT auth + role-based admin dashboards

### Quick Start

**Prerequisites**: Go 1.22+, PostgreSQL 14+

```bash
# 1. Set up the database
make setup-db

# 2. Run (migrations + default-account seeding happen automatically)
make dev
# or: go run cmd/server/main.go
```

<details>
<summary>Manual database setup (without <code>make</code>)</summary>

```bash
psql -U postgres << 'EOF'
CREATE USER altradits WITH PASSWORD 'password';
CREATE DATABASE altradits OWNER altradits;
GRANT ALL PRIVILEGES ON DATABASE altradits TO altradits;
EOF
```

```bash
DB_URL=postgres://altradits:password@localhost:5432/altradits?sslmode=disable go run cmd/server/main.go
```

</details>

**Default accounts**

| Role   | Email                    | Password   |
|--------|--------------------------|------------|
| Admin  | admin@altradits.com      | admin123   |
| Trader | trader@altradits.com     | trader123  |

### Architecture

- **Backend**: Go 1.22, standard library only. Custom PostgreSQL wire protocol driver (`internal/pgdrv`).
- **Frontend**: Pure HTML + CSS + vanilla JS. No frameworks.
- **Database**: PostgreSQL 14+
- **Auth**: SHA-256 session tokens, HttpOnly cookies

<details>
<summary>Routes</summary>

| Path | Role | Description |
|------|------|-------------|
| `/` | Public | Landing page |
| `/register` | Public | Create account |
| `/login` | Public | Sign in |
| `/customer/dashboard` | Customer | Wallet overview |
| `/customer/deposit` | Customer | M-Pesa deposit |
| `/customer/withdraw` | Customer | M-Pesa withdrawal |
| `/customer/investments` | Customer | Locked sats |
| `/admin/dashboard` | Admin | Operations overview |
| `/admin/deposits` | Admin | Approve deposits |
| `/trader/dashboard` | Trader | Pool management |

</details>

### Values, from the Bitcoin Whitepaper

1. **Trust minimization** — no ngurus, public assets
2. **Patience as proof of work** — 5-year lock minimum
3. **Timestamped transparency** — every sat traced
4. **No KYC** — email or phone only
5. **Community review** — no central authority
6. **Incentive alignment** — 2% of profit only
7. **Simplicity** — 6 buttons max
8. **Privacy by default** — no name required
9. **Long-term horizon** — shortcuts are traps
10. **Open participation** — no minimums

### Roadmap

- **Now**: Manual M-Pesa + Lightning (admin-approved)
- **Phase 2**: M-Pesa API (Daraja) + LND integration
- **Phase 3**: Events, Hackathons, Travel, Crowdfunding modules

### Contributing

1. Fork the repo, then `git clone` your fork
2. Follow [Quick Start](#quick-start) above to get it running locally
3. Create a branch, make your change, and open a PR

Bug fixes, tests, and roadmap items above are all welcome — open an issue first for bigger changes.

<img src="https://capsule-render.vercel.app/api?type=rect&color=FF6B6B&height=3&width=100&section=footer" width="100%"/>

## 📈 Build Log

| Stage | Shipped | Skills |
|---|---|---|
| Foundations | Go modules, Docker, multi-schema Postgres | Tooling, schema design |
| Money Rails | Lightning wallet + M-Pesa STK, BTC price tracking | LND REST, Redis, workers |
| Growth Features | Auto-save, bills, net worth, investing, planner | Service layers, ledger math |
| Engagement | AI coach, companion, notifications, hackathon mode | Pipelines, rapid iteration |
| Now | Lightning addresses, treasury pool, admin dashboards | LNURL-pay, double-entry, RBAC |

<div align="center">

**🔥 Still showing up, every day**

<img src="https://streak-stats.demolab.com/?user=altradits&theme=tokyonight&hide_border=true&background=0D1117&ring=FF6B6B&fire=FF6B6B&currStreakLabel=274C77&sideLabels=c9d1d9&dates=8b949e" width="100%" alt="streak"/>

</div>

<img src="https://capsule-render.vercel.app/api?type=rect&color=274C77&height=3&width=100&section=footer" width="100%"/>

## 🛣️ Next Up

- 🔌 Real LND node (currently mock-Lightning for dev)
- 🧪 Test coverage for `internal/wallet` & `internal/treasury`
- 📡 Public Lightning address service (`*@altradits.com`)
- 🤝 Sats to cash agent network with women-led businesses near me

<img src="https://capsule-render.vercel.app/api?type=rect&color=FF6B6B&height=3&width=100&section=footer" width="100%"/>

## 📡 Open Channels

Hiring, or building something with Go + Lightning rails? Let's talk.

<div align="center">

[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/stanmobitech)
[![Email](https://img.shields.io/badge/Email-Say_Hello-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:altradits@gmail.com)
[![WhatsApp](https://img.shields.io/badge/WhatsApp-%2B254707172370-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)](https://wa.me/254707172370)
[![GitHub](https://img.shields.io/badge/GitHub-altradits-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/altradits)

<br/>

<img src="https://github-readme-activity-graph.vercel.app/graph?username=altradits&theme=tokyo-night&hide_border=true&bg_color=0D1117&color=274C77&line=FF6B6B&point=ffffff&area=true&area_color=FF6B6B" width="100%" alt="activity"/>

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:FF6B6B,100:274C77&height=120&section=footer" width="100%" alt="footer wave"/>

</div>
