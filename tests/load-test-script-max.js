import http from 'k6/http';
import { check } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export const options = {
    scenarios: {
      get_profiles: {
        executor: 'per-vu-iterations',
        vus: 2000,
        iterations: 300,
        maxDuration: '5m',
      },
    },
  };

export default function () {
  const res = http.get('http://localhost:1323/profiles?lastName=a&offset=0&limit=100');
  check(res, {'is status 200': (r) => r.status === 200,});
}

export function handleSummary(data) {
    return {
        "summary.html": htmlReport(data),
    };
}