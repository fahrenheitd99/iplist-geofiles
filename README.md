# IPLIST GEOFILES

Automated build system for `geosite.dat` and `geoip.dat` for routing. Auto-updated every 2 days.

---

### 📥 Download Links

| File | GitHub Raw | jsDelivr CDN |
| :--- | :--- | :--- |
| **`geosite.dat`** | https://raw.githubusercontent.com/fahrenheitd99/iplist-geofiles/main/geosite.dat | https://cdn.jsdelivr.net/gh/fahrenheitd99/iplist-geofiles@main/geosite.dat |
| **`geoip.dat`** | https://raw.githubusercontent.com/fahrenheitd99/iplist-geofiles/main/geoip.dat | https://cdn.jsdelivr.net/gh/fahrenheitd99/iplist-geofiles@main/geoip.dat |

---

### 🏷️ Available Categories

| | | |
| :--- | :--- | :--- |
| `AI` | `HOSTING` | `SOCIALS` |
| `ANIME` | `JETBRAINS` | `TIKTOK` |
| `ART` | `MESSENGERS` | `TOOLS` |
| `CASINO` | `MUSIC` | `TORRENT` |
| `DISCORD` | `NEWS` | `VIDEO` |
| `EDUCATION` | `PORN` | `VK` |
| `FINANCE` | `PRIVATE` | `YANDEX` |
| `GAMES` | `RUSSIA` | `YOUTUBE` |
| `SHOP` | | |

---

<details>
<summary><b>Ruleset for v2rayNG and v2rayN (Click to expand)</b></summary>

<br>

**Domain Strategy:** IPIfNonMatch

**Rules source:** https://raw.githubusercontent.com/fahrenheitd99/iplist-geofiles/main/rules/v2rayn-rules.json

**Raw JSON**:
```json
[{"enabled":true,"locked":false,"outboundTag":"direct","protocol":["bittorrent"],"remarks":"Torrents"},{"enabled":true,"ip":["geoip:private"],"locked":false,"outboundTag":"direct","remarks":"private IPs"},{"domain":["geosite:private"],"enabled":true,"locked":false,"outboundTag":"direct","remarks":"private Domains"},{"enabled":true,"locked":false,"network":"udp","outboundTag":"proxy","port":"50000-65535","remarks":"voice2proxy"},{"domain":["geosite:ai","geosite:anime","geosite:art","geosite:casino","geosite:discord","geosite:education","geosite:finance","geosite:games","geosite:hosting","geosite:jetbrains","geosite:messengers","geosite:music","geosite:news","geosite:porn","geosite:shop","geosite:socials","geosite:tools","geosite:torrent","geosite:video","geosite:youtube"],"enabled":true,"locked":false,"outboundTag":"proxy","remarks":"Domains2proxy"},{"enabled":true,"ip":["geoip:ai","geoip:anime","geoip:art","geoip:casino","geoip:discord","geoip:education","geoip:finance","geoip:games","geoip:hosting","geoip:jetbrains","geoip:messengers","geoip:music","geoip:news","geoip:porn","geoip:shop","geoip:socials","geoip:tools","geoip:torrent","geoip:video","geoip:youtube"],"locked":false,"outboundTag":"proxy","remarks":"IPs2proxy"},{"enabled":true,"locked":false,"outboundTag":"direct","port":"0-65535","remarks":"Other"}]
```
</details>

**Data Sources:** https://iplist.opencck.org • https://beta.iplist.opencck.org • https://russia.iplist.opencck.org


