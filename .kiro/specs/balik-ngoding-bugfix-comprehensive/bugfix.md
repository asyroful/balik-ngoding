# Bugfix Requirements Document — Balik Ngoding Comprehensive

## Introduction

Balik Ngoding menghadapi tiga bug kritis yang mempengaruhi user experience, performance, dan reliability platform:

1. **Strict Answer Validation**: Validasi jawaban terlalu ketat, bahkan perbedaan spasi membuat jawaban dianggap salah, menyebabkan frustasi user pemula.
2. **Code Storage Best Practice**: Perlu review apakah menyimpan code soal di database adalah best practice atau sebaiknya di file system untuk performance dan maintainability.
3. **SQL Session Error 500**: Terjadi error 500 saat user mengerjakan soal SQL, mencegah user submit jawaban.

Spec ini menggunakan bug condition methodology untuk mendefinisikan setiap bug secara formal dengan acceptance criteria yang jelas.

---

## Bug #1: Strict Answer Validation

### Bug Analysis

#### Current Behavior (Defect)

1.1 WHEN user submits jawaban dengan output yang benar tetapi memiliki perbedaan whitespace (leading/trailing spaces, multiple spaces, newlines) THEN the system marks the answer as wrong_answer meskipun logika kode benar

1.2 WHEN user submits jawaban dengan output yang benar tetapi format berbeda (e.g., "hello world" vs "hello  world" dengan double space) THEN the system marks the answer as wrong_answer

1.3 WHEN user submits jawaban dengan output yang benar tetapi memiliki trailing newline berbeda THEN the system marks the answer as wrong_answer

#### Expected Behavior (Correct)

2.1 WHEN user submits jawaban dengan output yang benar tetapi memiliki perbedaan whitespace (leading/trailing spaces, multiple spaces, newlines) THEN the system SHALL normalize both actual dan expected output sebelum perbandingan dan mark the answer as accepted

2.2 WHEN user submits jawaban dengan output yang benar tetapi format berbeda (e.g., "hello world" vs "hello  world") THEN the system SHALL normalize whitespace dan mark the answer as accepted

2.3 WHEN user submits jawaban dengan output yang benar tetapi memiliki trailing newline berbeda THEN the system SHALL normalize newlines dan mark the answer as accepted

#### Unchanged Behavior (Regression Prevention)

3.1 WHEN user submits jawaban dengan output yang benar-benar berbeda (e.g., "hello" vs "goodbye") THEN the system SHALL CONTINUE TO mark the answer as wrong_answer

3.2 WHEN user submits jawaban dengan output yang benar untuk soal dengan multiple test cases THEN the system SHALL CONTINUE TO evaluate semua test cases dan return score yang akurat

3.3 WHEN user submits jawaban dengan runtime error THEN the system SHALL CONTINUE TO mark the answer as error dan return error message

---

## Bug #2: Code Storage Best Practice

### Bug Analysis

#### Current Behavior (Defect)

1.1 WHEN soal dibuat atau diupdate THEN the system stores problem code (description, starter code, schema) directly in PostgreSQL database sebagai text columns

1.2 WHEN user submits jawaban THEN the system stores user code directly in PostgreSQL database sebagai text column dalam Submission record

1.3 WHEN platform scales dengan ribuan soal dan submissions THEN the system experiences potential performance degradation karena large text blobs dalam database

#### Expected Behavior (Correct)

2.1 WHEN soal dibuat atau diupdate THEN the system SHALL evaluate apakah menyimpan code di database atau file system lebih optimal untuk performance, maintainability, dan scalability

2.2 WHEN code storage strategy ditentukan THEN the system SHALL implement best practice approach dengan dokumentasi alasan keputusan

2.3 WHEN user submits jawaban THEN the system SHALL store submission code menggunakan strategy yang telah ditentukan dengan proper indexing dan query optimization

#### Unchanged Behavior (Regression Prevention)

3.1 WHEN user retrieves problem detail THEN the system SHALL CONTINUE TO return complete problem information termasuk code dan schema

3.2 WHEN user retrieves submission history THEN the system SHALL CONTINUE TO return user code yang disubmit

3.3 WHEN analytics queries problem data THEN the system SHALL CONTINUE TO provide accurate statistics tanpa performance degradation

---

## Bug #3: SQL Session Error 500

### Bug Analysis

#### Current Behavior (Defect)

1.1 WHEN user submits SQL query untuk soal SQL THEN the system occasionally returns HTTP 500 error dengan message "Gagal memproses submission"

1.2 WHEN user submits SQL query dengan valid syntax THEN the system sometimes fails dengan error 500 instead of returning evaluation result

1.3 WHEN multiple users submit SQL queries simultaneously THEN the system may experience session/connection issues yang menghasilkan error 500

#### Expected Behavior (Correct)

2.1 WHEN user submits SQL query untuk soal SQL THEN the system SHALL successfully evaluate query dan return evaluation result dengan status accepted, wrong_answer, atau error

2.2 WHEN user submits SQL query dengan valid syntax THEN the system SHALL return proper evaluation result tanpa HTTP 500 error

2.3 WHEN multiple users submit SQL queries simultaneously THEN the system SHALL handle concurrent requests properly tanpa session/connection issues

#### Unchanged Behavior (Regression Prevention)

3.1 WHEN user submits SQL query dengan blocked operations (DROP, DELETE, UPDATE, INSERT, CREATE, ALTER, TRUNCATE) THEN the system SHALL CONTINUE TO reject query dengan error message "Operasi tidak diizinkan"

3.2 WHEN user submits SQL query dengan timeout THEN the system SHALL CONTINUE TO return error message "Waktu eksekusi habis"

3.3 WHEN user submits SQL query dengan syntax error THEN the system SHALL CONTINUE TO return error message dengan detail error

---

## Bug Condition Methodology

### Bug #1: Strict Answer Validation

**Bug Condition Function:**
```pascal
FUNCTION isBugCondition_StrictValidation(submission)
  INPUT: submission of type Submission
  OUTPUT: boolean
  
  RETURN (submission.actual_output != submission.expected_output) AND
         (normalize(submission.actual_output) == normalize(submission.expected_output))
END FUNCTION
```

**Property Specification - Fix Checking:**
```pascal
// Property: Whitespace Normalization
FOR ALL submission WHERE isBugCondition_StrictValidation(submission) DO
  result ← evaluator'(submission)
  ASSERT result.passed == true
END FOR
```

**Property Specification - Preservation Checking:**
```pascal
// Property: Preserve Wrong Answers
FOR ALL submission WHERE NOT isBugCondition_StrictValidation(submission) AND
                         normalize(submission.actual_output) != normalize(submission.expected_output) DO
  ASSERT evaluator'(submission).passed == false
END FOR
```

### Bug #2: Code Storage Best Practice

**Analysis Scope:**
- Review current database schema untuk Problem dan Submission models
- Evaluate performance impact dari storing large text blobs
- Compare database vs file system storage approaches
- Document recommendation dengan trade-offs

### Bug #3: SQL Session Error 500

**Bug Condition Function:**
```pascal
FUNCTION isBugCondition_SQLError500(submission)
  INPUT: submission of type Submission
  OUTPUT: boolean
  
  RETURN (submission.language == "sql") AND
         (submission.query_syntax_valid == true) AND
         (http_response_code == 500)
END FUNCTION
```

**Property Specification - Fix Checking:**
```pascal
// Property: SQL Submission Success
FOR ALL submission WHERE isBugCondition_SQLError500(submission) DO
  result ← submitSQL'(submission)
  ASSERT result.http_status_code == 200 AND
         result.body.status IN ["accepted", "wrong_answer", "error"]
END FOR
```

**Property Specification - Preservation Checking:**
```pascal
// Property: Preserve SQL Security
FOR ALL submission WHERE submission.language == "sql" AND
                         contains_blocked_operation(submission.query) DO
  result ← submitSQL'(submission)
  ASSERT result.body.error CONTAINS "Operasi tidak diizinkan"
END FOR
```

