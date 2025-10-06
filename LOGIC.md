# 🏛️ Sentiric Vertical Public Service - Mantık ve Akış Mimarisi

**Stratejik Rol:** Kamu hizmetleri, vergi veya kimlik doğrulama gibi devlet kurumlarına özgü süreçleri basitleştirir.

---

## 1. Temel Akış: Başvuru Gönderme (SubmitApplication)

```mermaid
sequenceDiagram
    participant Agent as Agent Service
    participant VPS as Public Service
    participant GovAPI as Harici Kamu API'si
    
    Agent->>VPS: SubmitApplication(type, form_data, user_id)
    
    Note over VPS: 1. Gerekli Formalite ve Veri Dönüşümü
    VPS->>GovAPI: POST /application (JSON/XML)
    GovAPI-->>VPS: Başvuru Takip ID'si
    
    Note over VPS: 2. Başarı Durumu ve Geri Bildirim
    VPS-->>Agent: SubmitApplicationResponse(tracking_id, message)
```

## 2. Hassasiyet ve Adaptasyon
Kamu hizmetleri, genellikle XML veya SOAP gibi eski protokoller kullanır ve yüksek güvenlik (e-devlet entegrasyonları) gerektirir. Bu servis, bu karmaşık protokolleri modern gRPC arayüzüne çeviren adaptör katmanı olarak işlev görür.