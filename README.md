# Uptime Monitor

<!-- 1–2 Sätze: Was macht das Projekt und warum gibt es es? -->

Ein selbst gehosteter Uptime-Monitor in Go. Er prüft regelmäßig Websites, speichert die Ergebnisse in PostgreSQL und zeigt sie über eine REST-API und eine Statusseite an.

> 🚧 Work in progress – siehe [Roadmap](#roadmap)

<!-- Später: CI-Badge, Screenshot der Statusseite -->

## Features

<!-- Nur eintragen, was wirklich funktioniert. Abhaken, sobald fertig. -->

- [ ] Parallele HTTP-Checks mit Timeout
- [ ] Speicherung in PostgreSQL
- [ ] REST-API zum Verwalten der Monitore
- [ ] Statusseite
- [ ] Docker-Setup
- [ ] Prometheus-Metriken & Grafana-Dashboard

## Tech-Stack

<!-- Ergänzen, sobald es im Projekt ist -->

- Go (Standardbibliothek `net/http`)

## Architektur

<!-- Später: Diagramm (z. B. Mermaid) + 2–3 Sätze, wie die Teile zusammenspielen -->

## Setup

### Voraussetzungen

- Go 1.xx <!-- deine Version: go version -->

### Starten

```bash
git clone https://github.com/Symeks/uptime-monitor.git
cd uptime-monitor
go run ./cmd/uptime-monitor
```

## API

<!-- Phase 3: Tabelle mit Methode | Pfad | Beschreibung + ein Beispiel-Request mit curl -->

## Design-Entscheidungen

<!-- Warum hast du etwas so gelöst? Kurz begründen. -->

## Roadmap

- [x] Phase 0 – Projekt-Setup
- [ ] Phase 1 – Checker
- [ ] Phase 2 – Datenbank
- [ ] Phase 3 – API & Statusseite
- [ ] Phase 4 – Docker
- [ ] Phase 5 – Logging & Observability
- [ ] Phase 6 – CI/CD
- [ ] Phase 7 – Deployment (AWS)
- [ ] Phase 8 – Kubernetes
- [ ] Phase 9 – Security

## Was ich gelernt habe

<!-- Am Ende: 3–5 ehrliche Punkte -->

## Lizenz

MIT
