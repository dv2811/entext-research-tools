package substack

// unimplemenented ideas
// curl 'https://substack.com/api/v1/realtime/token?channels=user%3A30913' \
//   -H 'accept: */*' \
//   -H 'accept-language: en-US,en;q=0.9' \
//   -H 'cache-control: max-age=0' \
//   -b 'ab_experiment_sampled=%22false%22; ab_testing_id=%22697ee87f-6774-49b5-8566-04ad290b581a%22; cookie_storage_key=760c86ac-f569-4bb3-99e2-6a4cb8f28da5; _ga=GA1.1.1138414561.1769069526; _gcl_au=1.1.1484436806.1769069525.1147152376.1769071112.1769071114; substack.sid=s%3AcWi5e7LdhiehyVFz6jXqMekbRg6exXJm.HzeK5fB7cKmTWkg9pvTxAJrj9x4dsTaCqcw4RJYy3VY; ajs_anonymous_id=%225178a3322f3721622ad001f827f119b9%22; muxData=mux_viewer_id=b5a736f2-3520-4d2c-9388-98177f012a0f&msn=0.38817064781826505&sid=2a7e2b5f-76fe-49f2-9ab8-270550c6e91b&sst=1769161016077&sex=1769162516079; _ga_TLW0DF6G5V=GS2.1.s1769161006$o3$g1$t1769161108$j60$l0$h0; visit_id=%7B%22id%22%3A%22087676ea-4f54-4a13-8722-b39c429799b9%22%2C%22timestamp%22%3A%222026-02-06T09%3A00%3A04.943Z%22%7D; substack.lli=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMwOTEzLCJpYXQiOjE3NzAzNjg0MDUsImV4cCI6MTc3Mjk2MDQwNSwiYXVkIjoibGlrZWx5LWxvZ2dlZC1pbiJ9.C_mvpTi--o-S8Qo_O9bJgo1l1a30XeeRt2mrGF56fHk; __cf_bm=X4lRAxLFu8T5RDg5Rva7tajjJSBrbA33YbMbyrEjhC0-1770368405-1.0.1.1-HGFgUgN8z9t03EbXeTZMiLv8HIGr9RuVBBZQBcSkt237KFCrYL5LItVN38i.iLRtJDdL_TWxPEUht9iTxFpg3j1Awz66u5SdDdQmq0EzYIM; _dd_s=rum=0&expire=1770369307105; cf_clearance=TJHHz_bfyzO5qS2F6RY4dJ7QcNoFxWLy0DHq9T77Rr8-1770368407-1.2.1.1-mzFpj0M0XSf_k_IhzIVb0G3DW1iHJ1SSHdyUg4mWTI6MfoLZpwL7UE16ujiKNtNV2SjHq8AmonR.eoM0bwGvpTtcKgQmdz7cJK_MrkOMdHmTJFXoVTxkyiVnFoQzv6I0zLeQVdIgjY64EQevHj8Nq6pvne9D9dILEF8rzUqyMXx3S0skvZIA.33dI9Ry2L4FbyWoe7yXGvub1CQIaDF5ZIcMsQ.gg9IQgzcLZDQyao8; AWSALBTG=v3ui2l1Wh4mm3OhgCfnYIqaQT0dEQBFGQj2SFK0pV+ujW+S4oD+GCJ2FrN+EPggbcuXWKGUMyP+GrbZPQhlXT3fenqNbq3sZCbqs/dcLGTi9sMBxzVrmRZzds62hD5aBHwYRKWhIM6jVp9+GGq+cLBS9Dj/PQbTQepotWHpHx2Wf; AWSALBTGCORS=v3ui2l1Wh4mm3OhgCfnYIqaQT0dEQBFGQj2SFK0pV+ujW+S4oD+GCJ2FrN+EPggbcuXWKGUMyP+GrbZPQhlXT3fenqNbq3sZCbqs/dcLGTi9sMBxzVrmRZzds62hD5aBHwYRKWhIM6jVp9+GGq+cLBS9Dj/PQbTQepotWHpHx2Wf' \
//   -H 'if-none-match: W/"2e3-9TO8HpMaraz0RC56XMY45YLow4s"' \
//   -H 'priority: u=1, i' \
//   -H 'referer: https://substack.com/home' \
//   -H 'sec-ch-ua: "Not(A:Brand";v="8", "Chromium";v="144"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36'

// response
// {
//     "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLTMwOTEzIiwiaXNzIjoicHJvZC1pc3N1ZXIiLCJhdWQiOiJ6eW5jIiwiZXhwIjoxNzcwMzcyMDA3LCJpYXQiOjE3NzAzNjg0MDcsInBlcm1pc3Npb25zIjpbInN1YnNjcmliZXx1c2VyOjMwOTEzIiwidW5zdWJzY3JpYmV8dXNlcjozMDkxMyJdfQ.MF3NeHF42XQqhH6yxS6dUNN34ALp85PP6EAkO4vX_cUkWXTd-NMeDGKf4b30X_nyd7ViMcXag8fUr6wjhfQLNYoBQ2GcpIhjPQp8IWSlCRyx7g3hE5V8Jk0O4R3UQncbPf3BWB2RO0FTw3hGVqwUTjLS4-whJ8JXleO7gd4uhR1uIrCDXYUADRt0Zdtb2CG1NOFwbR7r9w_bM11GZWCO9MAfO0N35kRq0NurKVe5e0ieCgkEdFzHEdIkCt1Vt9f-ZfllZbH4EKetViL26XQGdSTqeklustzH14dB-LLI0WKVTUwpLoFJrkT8s5JqToRHxNjCNjPSf8R88KPJ4qvEvg",
//     "expiry": "2026-02-06T10:00:07.000Z",
//     "permissions": [
//         "subscribe|user:30913",
//         "unsubscribe|user:30913"
//     ],
//     "endpoint": "wss://zyncrealtime.substack.com"
// }
