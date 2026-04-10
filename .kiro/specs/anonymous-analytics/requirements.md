# Requirements Document

## Introduction

Fitur Anonymous Analytics memungkinkan platform Balik Ngoding untuk melacak aktivitas pengguna tanpa sistem autentikasi. Setiap perangkat diidentifikasi menggunakan UUID anonim yang disimpan di localStorage. ID ini dikirim bersama setiap submission sehingga backend dapat menghitung estimasi jumlah pengguna unik, soal paling banyak dicoba, dan statistik agregat lainnya. Data ini ditampilkan di halaman dashboard statistik sederhana.

## Glossary

- **Anonymous_ID**: UUID v4 yang di-generate di sisi klien dan disimpan di localStorage sebagai identitas perangkat anonim.
- **Analytics_Service**: Komponen backend yang menghitung dan menyajikan statistik agregat dari data submissions.
- **Analytics_Dashboard**: Halaman frontend yang menampilkan statistik agregat platform.
- **ID_Generator**: Fungsi frontend yang membuat dan mempertahankan Anonymous_ID di localStorage.
- **Submission**: Data jawaban yang dikirim user ke backend, mencakup kode, bahasa, dan problem ID.
- **Unique_Device**: Perangkat yang diidentifikasi oleh satu Anonymous_ID unik.

---

## Requirements

### Requirement 1: Anonymous ID Generation & Persistence

**User Story:** As a platform visitor, I want my device to be identified by a persistent anonymous ID, so that my activity can be tracked across sessions without requiring an account.

#### Acceptance Criteria

1. WHEN a user visits the platform for the first time, THE ID_Generator SHALL create a UUID v4 and store it in localStorage under the key `balik-ngoding-anonymous-id`.
2. WHEN a user revisits the platform, THE ID_Generator SHALL read the existing Anonymous_ID from localStorage instead of generating a new one.
3. IF localStorage is unavailable (e.g., private browsing with storage blocked), THEN THE ID_Generator SHALL generate a session-scoped Anonymous_ID in memory without persisting it.
4. THE ID_Generator SHALL produce Anonymous_IDs that conform to the UUID v4 format (`xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx`).

---

### Requirement 2: Anonymous ID Submission

**User Story:** As a platform visitor, I want my anonymous ID to be sent with every submission, so that the platform can associate my activity with my device.

#### Acceptance Criteria

1. WHEN a user submits a solution, THE Submission SHALL include the Anonymous_ID as the field `anonymousId` in the request body sent to `POST /submit`.
2. WHEN the backend receives a submission with a valid `anonymousId`, THE Analytics_Service SHALL persist the `anonymousId` alongside the submission record.
3. IF a submission is received without an `anonymousId`, THEN THE Analytics_Service SHALL accept the submission and store a null value for `anonymousId`.
4. THE Submission SHALL NOT include any personally identifiable information beyond the Anonymous_ID.

---

### Requirement 3: Backend Analytics Aggregation

**User Story:** As a platform owner, I want to query aggregated statistics from the backend, so that I can understand platform usage without accessing raw user data.

#### Acceptance Criteria

1. WHEN `GET /analytics/stats` is called, THE Analytics_Service SHALL return the total count of distinct `anonymousId` values as `uniqueDevices`.
2. WHEN `GET /analytics/stats` is called, THE Analytics_Service SHALL return the total count of all submissions as `totalSubmissions`.
3. WHEN `GET /analytics/stats` is called, THE Analytics_Service SHALL return the total count of submissions with `status = 'accepted'` as `totalAccepted`.
4. WHEN `GET /analytics/stats` is called, THE Analytics_Service SHALL return a list of the top 5 problems ordered by submission count descending, each containing `problemId`, `title`, and `submissionCount`, as `topProblems`.
5. WHEN `GET /analytics/stats` is called, THE Analytics_Service SHALL return the count of distinct `anonymousId` values that have at least one submission with `status = 'accepted'` as `devicesWithAccepted`.
6. IF the submissions table is empty, THEN THE Analytics_Service SHALL return zero values for all numeric fields and an empty array for `topProblems`.

---

### Requirement 4: Analytics Dashboard

**User Story:** As a platform owner, I want a simple dashboard page to view platform statistics, so that I can monitor usage at a glance.

#### Acceptance Criteria

1. WHEN a user navigates to `/dashboard`, THE Analytics_Dashboard SHALL fetch and display the aggregated statistics from `GET /analytics/stats`.
2. WHEN the statistics are loading, THE Analytics_Dashboard SHALL display a loading skeleton in place of each statistic card.
3. IF the fetch request fails, THEN THE Analytics_Dashboard SHALL display an error message indicating that statistics could not be loaded.
4. THE Analytics_Dashboard SHALL display `uniqueDevices`, `totalSubmissions`, `totalAccepted`, `devicesWithAccepted`, and the `topProblems` list.
5. THE Analytics_Dashboard SHALL NOT display any raw Anonymous_IDs or individual submission records.

---

### Requirement 5: Data Integrity & Round-Trip

**User Story:** As a developer, I want the Anonymous_ID to be consistently serialized and deserialized, so that the same device is always identified correctly.

#### Acceptance Criteria

1. THE ID_Generator SHALL produce an Anonymous_ID such that reading it back from localStorage returns a string equal to the originally stored value (round-trip property).
2. WHEN an Anonymous_ID is included in a submission request body, THE Analytics_Service SHALL store the value without modification.
3. FOR ALL valid Anonymous_IDs stored in the database, querying the distinct count SHALL count each unique value exactly once (idempotence property).
