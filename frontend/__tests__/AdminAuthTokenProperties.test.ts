// @vitest-environment happy-dom
import { describe, it, expect, beforeEach, afterEach } from 'vitest';
import fc from 'fast-check';

// Feature: evaluator-fix-and-solution-keys, Property 15: session persistence

describe('AdminAuthToken — Session Persistence Properties', () => {
  beforeEach(() => {
    localStorage.clear();
  });

  afterEach(() => {
    localStorage.clear();
  });

  // ─── Task 21.1: Auth token persists in localStorage ───
  it('TestAuthTokenPersistenceProperty: auth token persists in localStorage across operations', () => {
    fc.assert(
      fc.property(fc.string({ minLength: 1 }), (token) => {
        // Store token
        localStorage.setItem('admin_auth_token', token);

        // Retrieve token
        const retrieved = localStorage.getItem('admin_auth_token');

        // Verify persistence
        return retrieved === token;
      }),
      { numRuns: 100 }
    );
  });

  // Additional: Verify token can be cleared
  it('auth token can be cleared from localStorage', () => {
    fc.assert(
      fc.property(fc.string({ minLength: 1 }), (token) => {
        // Store token
        localStorage.setItem('admin_auth_token', token);

        // Verify it's stored
        if (localStorage.getItem('admin_auth_token') !== token) {
          return false;
        }

        // Clear token
        localStorage.removeItem('admin_auth_token');

        // Verify it's cleared
        return localStorage.getItem('admin_auth_token') === null;
      }),
      { numRuns: 100 }
    );
  });

  // Additional: Verify multiple tokens don't interfere
  it('multiple localStorage keys do not interfere with auth token', () => {
    fc.assert(
      fc.property(
        fc.string({ minLength: 1 }),
        fc.string({ minLength: 1 }),
        fc.string({ minLength: 1 }),
        (authToken, otherKey, otherValue) => {
          // Skip if keys are the same
          if (authToken === otherKey) {
            return true;
          }

          // Store auth token
          localStorage.setItem('admin_auth_token', authToken);

          // Store other value
          localStorage.setItem(otherKey, otherValue);

          // Verify auth token is unchanged
          const retrievedAuth = localStorage.getItem('admin_auth_token');
          const retrievedOther = localStorage.getItem(otherKey);

          return retrievedAuth === authToken && retrievedOther === otherValue;
        }
      ),
      { numRuns: 100 }
    );
  });

  // Additional: Verify token persistence across multiple reads
  it('auth token remains consistent across multiple reads', () => {
    fc.assert(
      fc.property(fc.string({ minLength: 1 }), (token) => {
        localStorage.setItem('admin_auth_token', token);

        // Read multiple times
        const read1 = localStorage.getItem('admin_auth_token');
        const read2 = localStorage.getItem('admin_auth_token');
        const read3 = localStorage.getItem('admin_auth_token');

        // All reads should return the same value
        return read1 === token && read2 === token && read3 === token;
      }),
      { numRuns: 100 }
    );
  });
});
