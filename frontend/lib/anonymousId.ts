export const ANONYMOUS_ID_KEY = 'balik-ngoding-anonymous-id';

const UUID_V4_REGEX = /^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i;

function generateUUIDv4(): string {
  if (typeof crypto !== 'undefined' && crypto.randomUUID) {
    return crypto.randomUUID();
  }
  // Fallback for environments without crypto.randomUUID
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
    const r = (Math.random() * 16) | 0;
    const v = c === 'x' ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

export function isValidUUIDv4(value: string): boolean {
  return UUID_V4_REGEX.test(value);
}

export function getOrCreateAnonymousId(): string {
  try {
    const existing = localStorage.getItem(ANONYMOUS_ID_KEY);
    if (existing && isValidUUIDv4(existing)) {
      return existing;
    }
    const id = generateUUIDv4();
    localStorage.setItem(ANONYMOUS_ID_KEY, id);
    return id;
  } catch {
    // localStorage unavailable (e.g. private browsing with storage blocked)
    return generateUUIDv4();
  }
}
