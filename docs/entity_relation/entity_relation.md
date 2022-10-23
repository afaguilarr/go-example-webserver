Entity-Relation diagram:
```mermaid
erDiagram
          MASTER ||--|{ FAMILY_MEMBERSHIP : "has"
          FAMILY_MEMBERSHIP }|--|| FAMILY : "defines"
          FAMILY ||--o{ PET_OWNERSHIP : has
          PET_PROFILE ||--|{ PET_OWNERSHIP : "is owned by"

          MASTER ||--o| USER_PROFILE : "has"
          USER_PROFILE ||--o{ FAMILY_MANAGEMENT : "manages"
          FAMILY ||--|{ FAMILY_MANAGEMENT : "is managed by"

          PET_PROFILE ||--o{ VACCINE : "has"
          PET_PROFILE ||--o{ DISEASE : "has"
          PET_PROFILE }|--|| BREED : "has"
          PET_PROFILE ||--o{ MEDICINE : "needs"

          USER_PROFILE ||--o{ SUSCRIPTION : "follows"
          PET_PROFILE ||--o{ SUSCRIPTION : "is followed by"

          PET_PROFILE ||--o{ POST : "has"
          USER_PROFILE ||--o{ POST : "manages"
          POST ||--o{ COMMENT : "has"
          COMMENT }|--o| USER_PROFILE : "is done by (one of)"
          COMMENT }|--o| PET_PROFILE : "is done by (one of)"

          FAMILY ||--o{ USER_INVITATION : "sends"
          USER_PROFILE ||--o{ USER_INVITATION : "receives"
          FAMILY ||--o{ USER_REQUEST : "receives"
          USER_PROFILE ||--o{ USER_REQUEST : "sends"

          FAMILY ||--o{ PET_REQUEST_MOVE : "sends"
          FAMILY ||--o{ PET_REQUEST_MOVE : "receives"
          FAMILY ||--o{ PET_REQUEST_JOIN : "receives"
          FAMILY ||--o{ PET_REQUEST_JOIN : "sends"

          MASTER {
            string id PK
            string contact_number "optional"
            string name
            location location "optional"
          }

          USER_PROFILE {
            string id PK
            string username "unique"
            string description "optional"
            image profile_pic "optional"
            string master_id FK "unique"
          }

          PET_PROFILE {
            string id PK
            string species
            gender gender
            string name
            string nickname "unique"
            string description "optional"
            image profile_pic "optional"
            string breed_id FK
            location location "optional"
          }

          FAMILY {
            string id PK
            string name
            string nickname "unique"
            string description "optional"
          }

          VACCINE {
            string id PK
            string name
            string disease
            string description "optional"
            string brand "optional"
            string pet_id FK
          }

          DISEASE {
            string id PK
            string name
            string description
            string pet_profile_id FK
          }

          BREED {
            string id PK
            string species
            string name
            string description
          }

          MEDICINE {
            string id PK
            string pet_profile_id FK
            string name
            string brand "optional"
            string disease
          }

          POST {
            string id PK
            string pet_profile_id FK
            string user_profile_id FK
            time time
            bool comments_allowed
            string description
            image pic "optional"
            location location "optional"
          }

          COMMENT {
            string id PK
            string user_profile_id FK
            string pet_profile_id FK
            string content
            image pic "optional"
          }

          USER_INVITATION {
            string id PK
            string inviting_user_profile_id FK
            string invited_user_profile_id FK
            string family_id FK
            string invitation_content "optional"
            bool approved "optional"
            string response_content "optional"
          }

          USER_REQUEST {
            string id PK
            string requesting_user_profile_id FK
            string reviewer_user_profile_id FK
            string family_id FK
            string request_content "optional"
            bool approved "optional"
            string response_content "optional"
          }

          PET_REQUEST_MOVE {
            string id PK
            string requester_user_profile_id FK
            string pet_profile_id FK
            string family_id FK
            string reviewer_user_profile_id FK
            string request_content "optional"
            bool approved "optional"
            string response_content "optional"
          }

          PET_REQUEST_JOIN {
            string id PK
            string requester_user_profile_id FK
            string pet_profile_id FK
            string family_id FK
            string reviewer_user_profile_id FK
            string request_content "optional"
            bool approved "optional"
            string response_content "optional"
          }

          FAMILY_MEMBERSHIP {
            string id PK
            string master_id FK
            string family_id FK
          }

          FAMILY_MANAGEMENT {
            string id PK
            string user_profile_id FK
            string family_id FK
          }

          PET_OWNERSHIP {
            string id PK
            string pet_profile_id FK
            string family_id FK
          }

          SUSCRIPTION {
            string id PK
            string user_profile_id FK
            string pet_profile_id FK
          }
```