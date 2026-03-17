package substack

// unimplemenented ideas
// example request
// curl 'https://substack.com/api/v1/reader/feed?tab=for-you&commentId=209255871' \
//   -H 'accept: */*' \
//   -H 'accept-language: en-US,en;q=0.9' \
//   -H 'cache-control: max-age=0' \
//   -b 'ab_experiment_sampled=%22false%22; ab_testing_id=%22697ee87f-6774-49b5-8566-04ad290b581a%22; cookie_storage_key=760c86ac-f569-4bb3-99e2-6a4cb8f28da5; _ga=GA1.1.1138414561.1769069526; _gcl_au=1.1.1484436806.1769069525.1147152376.1769071112.1769071114; substack.sid=s%3AcWi5e7LdhiehyVFz6jXqMekbRg6exXJm.HzeK5fB7cKmTWkg9pvTxAJrj9x4dsTaCqcw4RJYy3VY; ajs_anonymous_id=%225178a3322f3721622ad001f827f119b9%22; muxData=mux_viewer_id=b5a736f2-3520-4d2c-9388-98177f012a0f&msn=0.38817064781826505&sid=2a7e2b5f-76fe-49f2-9ab8-270550c6e91b&sst=1769161016077&sex=1769162516079; visit_id=%7B%22id%22%3A%22087676ea-4f54-4a13-8722-b39c429799b9%22%2C%22timestamp%22%3A%222026-02-06T09%3A00%3A04.943Z%22%7D; substack.lli=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMwOTEzLCJpYXQiOjE3NzAzNjg0MDUsImV4cCI6MTc3Mjk2MDQwNSwiYXVkIjoibGlrZWx5LWxvZ2dlZC1pbiJ9.C_mvpTi--o-S8Qo_O9bJgo1l1a30XeeRt2mrGF56fHk; __cf_bm=X4lRAxLFu8T5RDg5Rva7tajjJSBrbA33YbMbyrEjhC0-1770368405-1.0.1.1-HGFgUgN8z9t03EbXeTZMiLv8HIGr9RuVBBZQBcSkt237KFCrYL5LItVN38i.iLRtJDdL_TWxPEUht9iTxFpg3j1Awz66u5SdDdQmq0EzYIM; cf_clearance=TJHHz_bfyzO5qS2F6RY4dJ7QcNoFxWLy0DHq9T77Rr8-1770368407-1.2.1.1-mzFpj0M0XSf_k_IhzIVb0G3DW1iHJ1SSHdyUg4mWTI6MfoLZpwL7UE16ujiKNtNV2SjHq8AmonR.eoM0bwGvpTtcKgQmdz7cJK_MrkOMdHmTJFXoVTxkyiVnFoQzv6I0zLeQVdIgjY64EQevHj8Nq6pvne9D9dILEF8rzUqyMXx3S0skvZIA.33dI9Ry2L4FbyWoe7yXGvub1CQIaDF5ZIcMsQ.gg9IQgzcLZDQyao8; _ga_TLW0DF6G5V=GS2.1.s1770368407$o4$g0$t1770368407$j60$l0$h0; _dd_s=rum=0&expire=1770369357482; AWSALBTG=K2Teoh4Y9KBKpxG0TLBODuRlKq7EoPIRnRHOqqwXoJDFL9GglWLWAO6M3Tquqvy9O3uUG1SKr3/+q2mXB7PAg9xAuQLfDnrPjNf1/VXpbXdCa1Im4uvF6VuKw5dGZhYM1XkN0bXg97C4naN58NzGgU203FKqPXrSA7SmY94m0QBF; AWSALBTGCORS=K2Teoh4Y9KBKpxG0TLBODuRlKq7EoPIRnRHOqqwXoJDFL9GglWLWAO6M3Tquqvy9O3uUG1SKr3/+q2mXB7PAg9xAuQLfDnrPjNf1/VXpbXdCa1Im4uvF6VuKw5dGZhYM1XkN0bXg97C4naN58NzGgU203FKqPXrSA7SmY94m0QBF' \
//   -H 'priority: u=1, i' \
//   -H 'referer: https://substack.com/@aisnakeoil/note/c-209255871' \
//   -H 'sec-ch-ua: "Not(A:Brand";v="8", "Chromium";v="144"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36'

// example response
// {
//     "items": [
//         {
//             "entity_key": "c-198924049",
//             "type": "comment",
//             "context": {
//                 "type": "note",
//                 "typeBucket": "notes",
//                 "timestamp": "2026-01-12T20:53:23.378Z",
//                 "users": [
//                     {
//                         "id": 4104627,
//                         "name": "Matt Mandel",
//                         "handle": "mandel",
//                         "previous_name": null,
//                         "photo_url": "https://bucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com/public/images/95b410fc-7f5b-47ee-a260-a24bb87ee915_398x500.jpeg",
//                         "bio": "matthewmandel.com",
//                         "profile_set_up_at": "2021-06-04T21:27:32.745Z",
//                         "reader_installed_at": "2022-03-09T19:33:05.620Z",
//                         "bestseller_tier": null,
//                         "status": {
//                             "bestsellerTier": null,
//                             "subscriberTier": null,
//                             "leaderboard": null,
//                             "vip": false,
//                             "badge": null,
//                             "paidPublicationIds": [],
//                             "subscriber": null
//                         },
//                         "primary_publication": {
//                             "id": 506418,
//                             "subdomain": "mandel",
//                             "custom_domain_optional": false,
//                             "name": "Matt Mandel",
//                             "author_id": 4104627,
//                             "user_id": 4104627,
//                             "handles_enabled": false,
//                             "explicit": false,
//                             "is_personal_mode": false,
//                             "payments_state": "disabled",
//                             "pledges_enabled": false
//                         }
//                     }
//                 ],
//                 "fallbackReason": "//TODO: create fallback reasons",
//                 "fallbackUrl": null,
//                 "isFresh": true,
//                 "source": "model",
//                 "model_rank": 0,
//                 "model_score": 1,
//                 "page": 0,
//                 "page_rank": 0
//             },
//             "publication": null,
//             "post": null,
//             "comment": {
//                 "name": "Matt Mandel",
//                 "handle": "mandel",
//                 "photo_url": "https://bucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com/public/images/95b410fc-7f5b-47ee-a260-a24bb87ee915_398x500.jpeg",
//                 "id": 198924049,
//                 "body": "I increasingly have friends ask what philosophy they should read to help think about AI. I've come to realize that using LLMs is pushing all of us to more closely examine our philosophical assumptions. Seemingly esoteric questions like who has moral status now feel urgent: people want to know if they should say \"please\" and \"thank you\" to Chat!  \n\nBelow is a rough attempt at laying out what in philosophy is most relevant for AI. I split the list between foundational philosophical concepts for thinking about AI and more applied frameworks for how to build and use AI\n\n\n\nPart 1: Philosophical Concepts\n\n\n\n\n\nWhat is a mind?\n\n\n\n\n\nNagel, What is it like to be a bat\n\n\n\nChalmers, Facing up to the problem of consciousness\n\n\n\nPutnam, The Nature of Mental States\n\n\n\nSearle, Minds, Brains, and Programs\n\n\n\nWhat is agency?\n\n\n\n\n\nAnscombe, Intention\n\n\n\nDennett, The Intentional Stance\n\n\n\nBratman, Shared Cooperative Activity\n\n\n\nApplbaum, Legitimacy (group agents)\n\n\n\nWho has moral status?\n\n\n\n\n\nKant, The Metaphysics of Morals (indirect duties)\n\n\n\nSinger, Animal Liberation\n\n\n\nKorsgaard, Fellow Creatures\n\n\n\nWhat are reasons and values?\n\n\n\n\n\nHume, Enquiry Concerning the Principles of Morals\n\n\n\nParfit, On What Matters (Volume II, Part 6)\n\n\n\nMackie, Ethics: Inventing Right and Wrong\n\n\n\nKorsgaard, The Sources of Normativity \n\nPart 2: How should we build AI\n\n\n\n\n\nHow might AIs come to implement our values? \n\n\n\n\n\nKripke, Wittgenstein on Rules and Private Language\n\n\n\nBostrom, The Superintelligent Will\n\n\n\nPlato, The Republic (Book I)\n\n\n\nAnthropic, Emergent Misalignment\n\n\n\nHow should AI be involved in governance?\n\n\n\n\n\nApplbaum, Legitimacy (“domination by a stone”)\n\n\n\n@Séb Krier, Coasean Bargaining at Scale\n\n\n\nNewsom, Snell v United Specialty Insurance Company (concurrence)\n\n\n\n@Cass Sunstein and Vermeule, Law and Leviathan\n\n\n\nHow might AI impact human flourishing?\n\n\n\n\n\nAristotle, Nicomachean Ethics\n\n\n\nMill, On Liberty\n\n\n\nFrankfurt, Freedom of the Will and the Concept of a Person\n\n\n\n@Brendan McCord, Live by the Claude, Die by the Claude\n\n\n\nHow do we find meaning in a post-AGI world?\n\n\n\n\n\nCamus, The Myth of Sisyphus\n\n\n\nNozick, Philosophical Explanations (Philosophy and the Meaning of Life)\n\n\n\nBostrom, Deep Utopia\n\n\n\nMandel, The Top Three AGI-Proof Careers and What They Reveal About Our Humanity and The Future",
//                 "body_json": {
//                     "type": "doc",
//                     "attrs": {
//                         "schemaVersion": "v1"
//                     },
//                     "content": [
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "I increasingly have friends ask what philosophy they should read to help think about AI. I've come to realize that using LLMs is pushing all of us to more closely examine our philosophical assumptions. Seemingly esoteric questions like who has moral status now feel urgent: people want to know if they should say \"please\" and \"thank you\" to Chat!  "
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Below is a rough attempt at laying out what in philosophy is most relevant for AI. I split the list between foundational philosophical concepts for thinking about AI and more applied frameworks for how to build and use AI"
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Part 1: Philosophical Concepts"
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "bulletList",
//                             "content": [
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "What is a mind?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Nagel, What is it like to be a bat"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Chalmers, Facing up to the problem of consciousness"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Putnam, The Nature of Mental States"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Searle, Minds, Brains, and Programs"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 },
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "What is agency?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Anscombe, Intention"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Dennett, The Intentional Stance"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Bratman, Shared Cooperative Activity"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Applbaum, Legitimacy (group agents)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 },
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "Who has moral status?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Kant, The Metaphysics of Morals (indirect duties)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Singer, Animal Liberation"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Korsgaard, Fellow Creatures"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 },
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "What are reasons and values?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Hume, Enquiry Concerning the Principles of Morals"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Parfit, On What Matters (Volume II, Part 6)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Mackie, Ethics: Inventing Right and Wrong"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Korsgaard, The Sources of Normativity "
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Part 2: How should we build AI"
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "bulletList",
//                             "content": [
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "How might AIs come to implement our values? "
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Kripke, Wittgenstein on Rules and Private Language"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Bostrom, The Superintelligent Will"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Plato, The Republic (Book I)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Anthropic, Emergent Misalignment"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 },
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "How should AI be involved in governance?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Applbaum, Legitimacy (“domination by a stone”)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "substack_mention",
//                                                                     "attrs": {
//                                                                         "id": 837581,
//                                                                         "label": "Séb Krier",
//                                                                         "mentionType": "user",
//                                                                         "url": null
//                                                                     }
//                                                                 },
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": ", Coasean Bargaining at Scale"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Newsom, Snell v United Specialty Insurance Company (concurrence)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "substack_mention",
//                                                                     "attrs": {
//                                                                         "id": 637324,
//                                                                         "label": "Cass Sunstein",
//                                                                         "mentionType": "user",
//                                                                         "url": null
//                                                                     }
//                                                                 },
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": " and Vermeule, Law and Leviathan"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 },
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "How might AI impact human flourishing?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Aristotle, Nicomachean Ethics"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Mill, On Liberty"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Frankfurt, Freedom of the Will and the Concept of a Person"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "substack_mention",
//                                                                     "attrs": {
//                                                                         "id": 866604,
//                                                                         "label": "Brendan McCord",
//                                                                         "mentionType": "user",
//                                                                         "url": null
//                                                                     }
//                                                                 },
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": ", Live by the Claude, Die by the Claude"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 },
//                                 {
//                                     "type": "listItem",
//                                     "content": [
//                                         {
//                                             "type": "paragraph",
//                                             "content": [
//                                                 {
//                                                     "type": "text",
//                                                     "text": "How do we find meaning in a post-AGI world?"
//                                                 }
//                                             ]
//                                         },
//                                         {
//                                             "type": "bulletList",
//                                             "content": [
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Camus, The Myth of Sisyphus"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Nozick, Philosophical Explanations (Philosophy and the Meaning of Life)"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Bostrom, Deep Utopia"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 },
//                                                 {
//                                                     "type": "listItem",
//                                                     "content": [
//                                                         {
//                                                             "type": "paragraph",
//                                                             "content": [
//                                                                 {
//                                                                     "type": "text",
//                                                                     "text": "Mandel, The Top Three AGI-Proof Careers and What They Reveal About Our Humanity and The Future"
//                                                                 }
//                                                             ]
//                                                         }
//                                                     ]
//                                                 }
//                                             ]
//                                         }
//                                     ]
//                                 }
//                             ]
//                         }
//                     ]
//                 },
//                 "publication_id": null,
//                 "post_id": null,
//                 "user_id": 4104627,
//                 "type": "feed",
//                 "date": "2026-01-12T20:53:23.378Z",
//                 "edited_at": null,
//                 "ancestor_path": "",
//                 "reply_minimum_role": "everyone",
//                 "media_clip_id": null,
//                 "reaction_count": 97,
//                 "reactions": {
//                     "❤": 97
//                 },
//                 "restacks": 13,
//                 "restacked": false,
//                 "children_count": 15,
//                 "attachments": [],
//                 "user_bestseller_tier": null,
//                 "userStatus": {
//                     "bestsellerTier": null,
//                     "subscriberTier": null,
//                     "leaderboard": null,
//                     "vip": false,
//                     "badge": null,
//                     "paidPublicationIds": [],
//                     "subscriber": null
//                 },
//                 "user_primary_publication": {
//                     "id": 506418,
//                     "subdomain": "mandel",
//                     "custom_domain_optional": false,
//                     "name": "Matt Mandel",
//                     "author_id": 4104627,
//                     "user_id": 4104627,
//                     "handles_enabled": false,
//                     "explicit": false,
//                     "is_personal_mode": false,
//                     "payments_state": "disabled",
//                     "pledges_enabled": false
//                 }
//             },
//             "parentComments": [],
//             "canReply": true,
//             "isMuted": false,
//             "trackingParameters": {
//                 "item_primary_entity_key": "c-198924049",
//                 "item_entity_key": "c-198924049",
//                 "item_type": "comment",
//                 "item_comment_id": 198924049,
//                 "item_content_user_id": 4104627,
//                 "item_content_timestamp": "2026-01-12T20:53:23.378Z",
//                 "item_context_type": "note",
//                 "item_context_type_bucket": "notes",
//                 "item_context_timestamp": "2026-01-12T20:53:23.378Z",
//                 "item_context_user_id": 4104627,
//                 "item_context_user_ids": [
//                     4104627
//                 ],
//                 "item_can_reply": true,
//                 "item_is_fresh": true,
//                 "item_last_impression_at": null,
//                 "item_source": "model",
//                 "item_model_rank": 0,
//                 "item_model_score": 1,
//                 "source_comment_id": 209255871,
//                 "item_page": 0,
//                 "item_page_rank": 0,
//                 "tab_id": "for-you",
//                 "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//                 "impression_id": "e1c121ba-6830-49bd-bd7d-c3207b1d091d",
//                 "followed_user_count": 224,
//                 "subscribed_publication_count": 205,
//                 "is_following": false,
//                 "is_explicitly_subscribed": false,
//                 "note_velocity_factor": 1.078759961279,
//                 "note_delay_seconds": 141,
//                 "note_notes_per_hour": 2819.426967,
//                 "item_current_reaction_count": 97,
//                 "item_current_restack_count": 13,
//                 "item_current_reply_count": 15
//             },
//             "canShowFollowUpsell": false
//         },
//         {
//             "entity_key": "c-196576534",
//             "type": "comment",
//             "context": {
//                 "type": "comment_like",
//                 "typeBucket": "notes",
//                 "timestamp": "2026-01-07T19:19:50.470Z",
//                 "users": [
//                     {
//                         "id": 10472909,
//                         "name": "Nathan Lambert",
//                         "handle": "natolambert",
//                         "previous_name": null,
//                         "photo_url": "https://substackcdn.com/image/fetch/$s_!RihO!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F8fedcdfb-e137-4f6a-9089-a46add6c6242_500x500.jpeg",
//                         "bio": "ML researcher making sense of AI research, products, and the uncertain technological future. PhD from Berkeley AI. Experience at Meta, DeepMind, HuggingFace.",
//                         "profile_set_up_at": "2021-04-24T01:19:33.371Z",
//                         "reader_installed_at": "2022-03-09T17:52:30.690Z",
//                         "bestseller_tier": 100,
//                         "status": {
//                             "bestsellerTier": 100,
//                             "subscriberTier": 5,
//                             "leaderboard": {
//                                 "ranking": "paid",
//                                 "rank": 39,
//                                 "publicationName": "Interconnects AI",
//                                 "label": "Technology",
//                                 "categoryId": "4",
//                                 "publicationId": 48206
//                             },
//                             "vip": false,
//                             "badge": {
//                                 "type": "bestseller",
//                                 "tier": 100
//                             },
//                             "paidPublicationIds": [
//                                 883883,
//                                 1084918,
//                                 6349492,
//                                 1915042,
//                                 69345,
//                                 6027
//                             ],
//                             "subscriber": null
//                         },
//                         "primary_publication": {
//                             "id": 48206,
//                             "subdomain": "robotic",
//                             "custom_domain": "www.interconnects.ai",
//                             "custom_domain_optional": false,
//                             "name": "Interconnects AI",
//                             "logo_url": "https://substack-post-media.s3.amazonaws.com/public/images/c52e8097-8f3d-4f7e-808b-2f4ad37f3b52_720x720.png",
//                             "author_id": 10472909,
//                             "user_id": 10472909,
//                             "handles_enabled": false,
//                             "explicit": false,
//                             "is_personal_mode": false,
//                             "payments_state": "enabled",
//                             "pledges_enabled": false
//                         }
//                     }
//                 ],
//                 "fallbackReason": "//TODO: create fallback reasons",
//                 "fallbackUrl": null,
//                 "isFresh": true,
//                 "source": "model",
//                 "model_rank": 1,
//                 "model_score": 0.5,
//                 "page": 0,
//                 "page_rank": 1
//             },
//             "publication": null,
//             "post": null,
//             "comment": {
//                 "name": "Abi Olvera",
//                 "handle": "abio",
//                 "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/550e023c-2e8e-440f-91e8-6d32872d8d5f_1123x1125.png",
//                 "id": 196576534,
//                 "body": "I waited tables for six years at $2.13 an hour. Now I'm a policy researcher in DC. These two worlds talk about gig work completely differently.\n\nMedia says Uber is exploitative because it doesn't offer health insurance or guaranteed wages. But most jobs my family members were looking at (daycare, retail, construction) don't either. My family advised me toward waiting tables precisely because its lack of guaranteed wages meant I could make more from tips.\n\nWhat my community actually complained about was Uber’s requirement that your car needs to be newer. And the app takes a 60% cut on some rides.\n\nWhat my community loves is that speaking English isn't required. Uber and Lyft will consider people with criminal records.\n\nThe way gig work is discussed doesn’t reflect what I see. We judge new technologies by the standards of people with better choices than the people actually using them.",
//                 "body_json": {
//                     "type": "doc",
//                     "attrs": {
//                         "schemaVersion": "v1"
//                     },
//                     "content": [
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "I waited tables for six years at $2.13 an hour. Now I'm a policy researcher in DC. These two worlds talk about gig work completely differently."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Media says Uber is exploitative because it doesn't offer health insurance or guaranteed wages. But most jobs my family members were looking at (daycare, retail, construction) don't either. My family advised me toward waiting tables precisely because its lack of guaranteed wages meant I could make more from tips."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "What my community actually complained about was Uber’s requirement that your car needs to be newer. And the app takes a 60% cut on some rides."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "What my community loves is that speaking English isn't required. Uber and Lyft will consider people with criminal records."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "The way gig work is discussed doesn’t reflect what I see. We judge new technologies by the standards of people with better choices than the people actually using them."
//                                 }
//                             ]
//                         }
//                     ]
//                 },
//                 "publication_id": null,
//                 "post_id": null,
//                 "user_id": 349629,
//                 "type": "feed",
//                 "date": "2026-01-07T19:19:50.470Z",
//                 "edited_at": "2026-01-07T19:42:56.053Z",
//                 "ancestor_path": "",
//                 "reply_minimum_role": "everyone",
//                 "media_clip_id": null,
//                 "reaction_count": 659,
//                 "reactions": {
//                     "❤": 659
//                 },
//                 "restacks": 34,
//                 "restacked": false,
//                 "children_count": 18,
//                 "attachments": [
//                     {
//                         "id": "93049c2d-6b47-41e1-8bdd-a37d9dc597ec",
//                         "type": "post",
//                         "publication": {
//                             "apple_pay_disabled": false,
//                             "apex_domain": null,
//                             "author_id": 349629,
//                             "byline_images_enabled": true,
//                             "bylines_enabled": true,
//                             "chartable_token": null,
//                             "community_enabled": true,
//                             "copyright": "Abigail Olvera",
//                             "cover_photo_url": null,
//                             "created_at": "2022-11-16T02:38:49.439Z",
//                             "custom_domain_optional": false,
//                             "custom_domain": null,
//                             "default_comment_sort": "best_first",
//                             "default_coupon": null,
//                             "default_group_coupon": null,
//                             "default_show_guest_bios": true,
//                             "email_banner_url": null,
//                             "email_from_name": "Abi Olvera's Newsletter",
//                             "email_from": null,
//                             "embed_tracking_disabled": false,
//                             "explicit": false,
//                             "expose_paywall_content_to_search_engines": true,
//                             "fb_pixel_id": null,
//                             "fb_site_verification_token": null,
//                             "flagged_as_spam": false,
//                             "founding_subscription_benefits": [
//                                 "Help me hit 99 founding members to be self sustaining and fully devote myself to this work"
//                             ],
//                             "free_subscription_benefits": [
//                                 "Occasional public posts"
//                             ],
//                             "ga_pixel_id": null,
//                             "google_site_verification_token": null,
//                             "google_tag_manager_token": null,
//                             "hero_image": null,
//                             "hero_text": "Optimistic takes on why things cost so much, why institutions make mistakes, and what AI is actually changing.",
//                             "hide_intro_subtitle": null,
//                             "hide_intro_title": null,
//                             "hide_podcast_feed_link": false,
//                             "homepage_type": "magaziney",
//                             "id": 1194762,
//                             "image_thumbnails_always_enabled": false,
//                             "invite_only": false,
//                             "hide_podcast_from_pub_listings": false,
//                             "language": "en",
//                             "logo_url_wide": null,
//                             "logo_url": null,
//                             "minimum_group_size": 2,
//                             "moderation_enabled": true,
//                             "name": "Positive Sum",
//                             "paid_subscription_benefits": [
//                                 "Subscriber-only posts and full archive",
//                                 "Help me hit 400 supporters so I am self sustaining and can do more deep dives"
//                             ],
//                             "parsely_pixel_id": null,
//                             "chartbeat_domain": null,
//                             "payments_state": "enabled",
//                             "paywall_free_trial_enabled": true,
//                             "podcast_art_url": null,
//                             "paid_podcast_episode_art_url": null,
//                             "podcast_byline": null,
//                             "podcast_description": null,
//                             "podcast_enabled": false,
//                             "podcast_feed_url": null,
//                             "podcast_title": null,
//                             "post_preview_limit": null,
//                             "primary_user_id": 349629,
//                             "require_clickthrough": false,
//                             "show_pub_podcast_tab": false,
//                             "show_recs_on_homepage": true,
//                             "subdomain": "abio",
//                             "subscriber_invites": 0,
//                             "support_email": null,
//                             "theme_var_background_pop": "#67BDFC",
//                             "theme_var_color_links": false,
//                             "theme_var_cover_bg_color": null,
//                             "trial_end_override": null,
//                             "twitter_pixel_id": null,
//                             "type": "newsletter",
//                             "post_reaction_faces_enabled": true,
//                             "is_personal_mode": false,
//                             "plans": [
//                                 {
//                                     "id": "yearly80usd",
//                                     "object": "plan",
//                                     "active": true,
//                                     "aggregate_usage": null,
//                                     "amount": 8000,
//                                     "amount_decimal": "8000",
//                                     "billing_scheme": "per_unit",
//                                     "created": 1748540077,
//                                     "currency": "usd",
//                                     "interval": "year",
//                                     "interval_count": 1,
//                                     "livemode": true,
//                                     "metadata": {
//                                         "substack": "yes"
//                                     },
//                                     "meter": null,
//                                     "nickname": "$80 a year",
//                                     "product": "prod_SOyI6vIH9Lx4ng",
//                                     "tiers": null,
//                                     "tiers_mode": null,
//                                     "transform_usage": null,
//                                     "trial_period_days": null,
//                                     "usage_type": "licensed",
//                                     "currency_options": {
//                                         "aud": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 12000,
//                                             "unit_amount_decimal": "12000"
//                                         },
//                                         "brl": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 44500,
//                                             "unit_amount_decimal": "44500"
//                                         },
//                                         "cad": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 11000,
//                                             "unit_amount_decimal": "11000"
//                                         },
//                                         "chf": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 6500,
//                                             "unit_amount_decimal": "6500"
//                                         },
//                                         "dkk": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 51000,
//                                             "unit_amount_decimal": "51000"
//                                         },
//                                         "eur": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 7000,
//                                             "unit_amount_decimal": "7000"
//                                         },
//                                         "gbp": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 6000,
//                                             "unit_amount_decimal": "6000"
//                                         },
//                                         "mxn": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 144500,
//                                             "unit_amount_decimal": "144500"
//                                         },
//                                         "nok": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 81000,
//                                             "unit_amount_decimal": "81000"
//                                         },
//                                         "nzd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 14000,
//                                             "unit_amount_decimal": "14000"
//                                         },
//                                         "pln": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 29000,
//                                             "unit_amount_decimal": "29000"
//                                         },
//                                         "sek": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 74000,
//                                             "unit_amount_decimal": "74000"
//                                         },
//                                         "usd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 8000,
//                                             "unit_amount_decimal": "8000"
//                                         }
//                                     }
//                                 },
//                                 {
//                                     "id": "monthly8usd",
//                                     "object": "plan",
//                                     "active": true,
//                                     "aggregate_usage": null,
//                                     "amount": 800,
//                                     "amount_decimal": "800",
//                                     "billing_scheme": "per_unit",
//                                     "created": 1748540076,
//                                     "currency": "usd",
//                                     "interval": "month",
//                                     "interval_count": 1,
//                                     "livemode": true,
//                                     "metadata": {
//                                         "substack": "yes"
//                                     },
//                                     "meter": null,
//                                     "nickname": "$8 a month",
//                                     "product": "prod_SOyI5XfMAvvRkI",
//                                     "tiers": null,
//                                     "tiers_mode": null,
//                                     "transform_usage": null,
//                                     "trial_period_days": null,
//                                     "usage_type": "licensed",
//                                     "currency_options": {
//                                         "aud": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 1200,
//                                             "unit_amount_decimal": "1200"
//                                         },
//                                         "brl": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 4500,
//                                             "unit_amount_decimal": "4500"
//                                         },
//                                         "cad": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 1100,
//                                             "unit_amount_decimal": "1100"
//                                         },
//                                         "chf": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 700,
//                                             "unit_amount_decimal": "700"
//                                         },
//                                         "dkk": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 5500,
//                                             "unit_amount_decimal": "5500"
//                                         },
//                                         "eur": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 700,
//                                             "unit_amount_decimal": "700"
//                                         },
//                                         "gbp": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 600,
//                                             "unit_amount_decimal": "600"
//                                         },
//                                         "mxn": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 14500,
//                                             "unit_amount_decimal": "14500"
//                                         },
//                                         "nok": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 8500,
//                                             "unit_amount_decimal": "8500"
//                                         },
//                                         "nzd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 1400,
//                                             "unit_amount_decimal": "1400"
//                                         },
//                                         "pln": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 2900,
//                                             "unit_amount_decimal": "2900"
//                                         },
//                                         "sek": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 7500,
//                                             "unit_amount_decimal": "7500"
//                                         },
//                                         "usd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 800,
//                                             "unit_amount_decimal": "800"
//                                         }
//                                     }
//                                 },
//                                 {
//                                     "id": "founding30000usd",
//                                     "name": "founding30000usd",
//                                     "nickname": "founding30000usd",
//                                     "active": true,
//                                     "amount": 30000,
//                                     "currency": "usd",
//                                     "interval": "year",
//                                     "interval_count": 1,
//                                     "metadata": {
//                                         "substack": "yes",
//                                         "founding": "yes",
//                                         "no_coupons": "yes",
//                                         "short_description": "Founding Member",
//                                         "short_description_english": "Founding Member",
//                                         "minimum": "8000",
//                                         "minimum_local": {
//                                             "aud": 11500,
//                                             "brl": 42500,
//                                             "cad": 11000,
//                                             "chf": 6500,
//                                             "dkk": 51000,
//                                             "eur": 7000,
//                                             "gbp": 6000,
//                                             "mxn": 139500,
//                                             "nok": 78500,
//                                             "nzd": 13500,
//                                             "pln": 29000,
//                                             "sek": 72500,
//                                             "usd": 8000
//                                         }
//                                     },
//                                     "currency_options": {
//                                         "aud": {
//                                             "unit_amount": 43500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "brl": {
//                                             "unit_amount": 158000,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "cad": {
//                                             "unit_amount": 41500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "chf": {
//                                             "unit_amount": 23500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "dkk": {
//                                             "unit_amount": 190000,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "eur": {
//                                             "unit_amount": 25500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "gbp": {
//                                             "unit_amount": 22500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "mxn": {
//                                             "unit_amount": 522500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "nok": {
//                                             "unit_amount": 293500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "nzd": {
//                                             "unit_amount": 50500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "pln": {
//                                             "unit_amount": 107500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "sek": {
//                                             "unit_amount": 271500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "usd": {
//                                             "unit_amount": 30000,
//                                             "tax_behavior": "unspecified"
//                                         }
//                                     }
//                                 }
//                             ],
//                             "stripe_user_id": "acct_1RUAK1GPYkoKP8us",
//                             "stripe_country": "US",
//                             "stripe_publishable_key": "pk_live_51RUAK1GPYkoKP8usgrgmcdW5q5OsZo3U5J8lnUEu5gHJuB3JfugwgE4BXntS8emj2Ecpd3peJPrBNwm7a2DOkN9q00KTjhi1AU",
//                             "stripe_platform_account": "US",
//                             "automatic_tax_enabled": false,
//                             "author_name": "Abi Olvera",
//                             "author_handle": "abio",
//                             "author_photo_url": "https://substackcdn.com/image/fetch/$s_!yJm5!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F550e023c-2e8e-440f-91e8-6d32872d8d5f_1123x1125.png",
//                             "author_bio": "Optimistic takes on why things cost so much, why institutions make mistakes, and what AI is actually changing. Ex-diplomat. I lead DC Abundance. Golden Gate Institute Research Director. Emergent Ventures grantee. Views my own.",
//                             "has_custom_tos": false,
//                             "has_custom_privacy": false,
//                             "theme": {
//                                 "background_pop_color": null,
//                                 "web_bg_color": "#ffffff",
//                                 "cover_bg_color": null,
//                                 "publication_id": 1194762,
//                                 "color_links": null,
//                                 "font_preset_heading": null,
//                                 "font_preset_body": null,
//                                 "font_family_headings": null,
//                                 "font_family_body": null,
//                                 "font_family_ui": null,
//                                 "font_size_body_desktop": null,
//                                 "print_secondary": null,
//                                 "custom_css_web": null,
//                                 "custom_css_email": null,
//                                 "home_hero": "magaziney",
//                                 "home_posts": "list",
//                                 "home_show_top_posts": true,
//                                 "hide_images_from_list": false,
//                                 "home_hero_alignment": "left",
//                                 "home_hero_show_podcast_links": true,
//                                 "default_post_header_variant": null,
//                                 "custom_header": null,
//                                 "custom_footer": null,
//                                 "social_media_links": null,
//                                 "font_options": null,
//                                 "section_template": null
//                             },
//                             "threads_v2_settings": null,
//                             "default_group_coupon_percent_off": null,
//                             "pause_return_date": null,
//                             "has_posts": true,
//                             "has_recommendations": true,
//                             "first_post_date": "2023-01-06T00:15:29.597Z",
//                             "has_podcast": false,
//                             "has_free_podcast": false,
//                             "has_subscriber_only_podcast": false,
//                             "has_community_content": true,
//                             "rankingDetail": "Launched 3 years ago",
//                             "rankingDetailFreeIncluded": "Thousands of subscribers",
//                             "rankingDetailOrderOfMagnitude": 10,
//                             "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                             "rankingDetailFreeSubscriberCount": "Over 1,000 subscribers",
//                             "rankingDetailByLanguage": {
//                                 "da": {
//                                     "rankingDetail": "Lancering 3 år",
//                                     "rankingDetailFreeIncluded": "Tusindvis af abonnenter",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Over 1,000 abonnenter",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "de": {
//                                     "rankingDetail": "Vor vor 3 Jahren gelauncht",
//                                     "rankingDetailFreeIncluded": "Tausende von Abonnenten",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Über 1,000 Abonnenten",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "es": {
//                                     "rankingDetail": "Lanzado hace 3 años",
//                                     "rankingDetailFreeIncluded": "Miles de suscriptores",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Más de 1,000 suscriptores",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "fr": {
//                                     "rankingDetail": "Lancé il y a 3 années",
//                                     "rankingDetailFreeIncluded": "Des milliers d'abonnés",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Plus de 1,000 abonnés",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "ja": {
//                                     "rankingDetail": "開始日 3年前",
//                                     "rankingDetailFreeIncluded": "登録者数数千人",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "1,000 名以上の購読者",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "nb": {
//                                     "rankingDetail": "Lansert 3 år",
//                                     "rankingDetailFreeIncluded": "Tusenvis av abonnenter",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Over 1,000 abonnenter",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "nl": {
//                                     "rankingDetail": "Gelanceerd 3 jaar geleden",
//                                     "rankingDetailFreeIncluded": "Duizenden abonnees",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Meer dan 1,000 abonnees",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "pl": {
//                                     "rankingDetail": "Uruchomiono 3 lat temu",
//                                     "rankingDetailFreeIncluded": "Tysiące subskrybentów",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Ponad 1,000 subskrybentów",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "pt": {
//                                     "rankingDetail": "Lançado 3 anos",
//                                     "rankingDetailFreeIncluded": "Milhares de subscritores",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Mais de 1,000 subscritores",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "pt-br": {
//                                     "rankingDetail": "Lançado 3 anos",
//                                     "rankingDetailFreeIncluded": "Milhares de assinantes",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Mais de 1,000 assinantes",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "it": {
//                                     "rankingDetail": "Lanciato 3 anni",
//                                     "rankingDetailFreeIncluded": "Migliaia di abbonati",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Oltre 1,000 abbonati",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "tr": {
//                                     "rankingDetail": "3 yıl başlatıldı",
//                                     "rankingDetailFreeIncluded": "Binlerce abone",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "1,000'in üzerinde abone",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "sv": {
//                                     "rankingDetail": "Lanserad 3 år sedan",
//                                     "rankingDetailFreeIncluded": "Tusentals prenumeranter",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Över 1,000 prenumeranter",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 },
//                                 "en": {
//                                     "rankingDetail": "Launched 3 years ago",
//                                     "rankingDetailFreeIncluded": "Thousands of subscribers",
//                                     "rankingDetailOrderOfMagnitude": 10,
//                                     "rankingDetailFreeIncludedOrderOfMagnitude": 1000,
//                                     "rankingDetailFreeSubscriberCount": "Over 1,000 subscribers",
//                                     "freeSubscriberCount": "1,000",
//                                     "freeSubscriberCountOrderOfMagnitude": "1K+"
//                                 }
//                             },
//                             "freeSubscriberCount": "1,000",
//                             "freeSubscriberCountOrderOfMagnitude": "1K+",
//                             "author_bestseller_tier": 0,
//                             "author_badge": {
//                                 "type": "subscriber",
//                                 "tier": 1,
//                                 "accent_colors": null
//                             },
//                             "disable_monthly_subscriptions": false,
//                             "disable_annual_subscriptions": false,
//                             "hide_post_restacks": false,
//                             "notes_feed_enabled": false,
//                             "showIntroModule": false,
//                             "isPortraitLayout": false,
//                             "last_chat_post_at": null,
//                             "primary_profile_name": "Abi Olvera",
//                             "primary_profile_photo_url": "https://substackcdn.com/image/fetch/$s_!yJm5!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F550e023c-2e8e-440f-91e8-6d32872d8d5f_1123x1125.png",
//                             "no_follow": false,
//                             "paywall_chat": "free",
//                             "sections": [],
//                             "multipub_migration": null,
//                             "navigationBarItems": [
//                                 {
//                                     "id": "1137971c-5b62-4760-8f9f-89f51cc762bb",
//                                     "publication_id": 1194762,
//                                     "sibling_rank": 0,
//                                     "link_title": "Abundance",
//                                     "link_url": "",
//                                     "section_id": null,
//                                     "post_id": null,
//                                     "is_hidden": null,
//                                     "standard_key": null,
//                                     "post_tag_id": "38871425-2a94-4f9e-b057-d901e70880b2",
//                                     "parent_id": null,
//                                     "is_group": false,
//                                     "post": null,
//                                     "section": null,
//                                     "postTag": {
//                                         "id": "38871425-2a94-4f9e-b057-d901e70880b2",
//                                         "publication_id": 1194762,
//                                         "name": "abundance",
//                                         "slug": "abundance",
//                                         "hidden": false
//                                     },
//                                     "children": []
//                                 },
//                                 {
//                                     "id": "a0121027-1ae2-4b7a-9d93-9f032aa59138",
//                                     "publication_id": 1194762,
//                                     "sibling_rank": 1,
//                                     "link_title": "Nuanced AI Takes",
//                                     "link_url": "",
//                                     "section_id": null,
//                                     "post_id": null,
//                                     "is_hidden": null,
//                                     "standard_key": null,
//                                     "post_tag_id": "f1f90931-4c21-4119-8f53-49f9f9c32bcb",
//                                     "parent_id": null,
//                                     "is_group": false,
//                                     "post": null,
//                                     "section": null,
//                                     "postTag": {
//                                         "id": "f1f90931-4c21-4119-8f53-49f9f9c32bcb",
//                                         "publication_id": 1194762,
//                                         "name": "AI",
//                                         "slug": "ai",
//                                         "hidden": false
//                                     },
//                                     "children": []
//                                 }
//                             ],
//                             "contributors": [
//                                 {
//                                     "name": "Abi Olvera",
//                                     "handle": "abio",
//                                     "role": "admin",
//                                     "owner": true,
//                                     "user_id": 349629,
//                                     "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/550e023c-2e8e-440f-91e8-6d32872d8d5f_1123x1125.png",
//                                     "bio": "Optimistic takes on why things cost so much, why institutions make mistakes, and what AI is actually changing. Ex-diplomat. I lead DC Abundance. Golden Gate Institute Research Director. Emergent Ventures grantee. Views my own."
//                                 }
//                             ],
//                             "threads_v2_enabled": false,
//                             "viralGiftsConfig": {
//                                 "id": "b98b974a-10e5-482a-af0f-567b81b8337e",
//                                 "publication_id": 1194762,
//                                 "enabled": true,
//                                 "gifts_per_user": 5,
//                                 "gift_length_months": 1,
//                                 "send_extra_gifts": true,
//                                 "message": "Making the pie bigger: innovation, resilience, and the tradeoffs that matter",
//                                 "created_at": "2025-05-29T17:54:08.14625+00:00",
//                                 "updated_at": "2025-05-29T17:54:08.14625+00:00",
//                                 "days_til_invite": 14,
//                                 "send_emails": true,
//                                 "show_link": null,
//                                 "grant_email_body": null,
//                                 "grant_email_subject": null
//                             },
//                             "tier": 2,
//                             "no_index": false,
//                             "can_set_google_site_verification": true,
//                             "can_have_sitemap": true,
//                             "draft_iap_advanced_plans": [
//                                 {
//                                     "sku": "VajhJQ5OQpwY5tB9Nh",
//                                     "publication_id": "1194762",
//                                     "is_active": true,
//                                     "price_base_units": 1100,
//                                     "currency_alpha3": "usd",
//                                     "period": "month",
//                                     "created_at": "2025-05-29T17:35:50.716Z",
//                                     "updated_at": "2025-05-29T17:35:50.716Z",
//                                     "id": "55247",
//                                     "payout_amount_base_units": 80,
//                                     "alternate_currencies": {
//                                         "aud": 1800,
//                                         "brl": 6500,
//                                         "cad": 1600,
//                                         "chf": 1000,
//                                         "dkk": 7500,
//                                         "eur": 1000,
//                                         "gbp": 900,
//                                         "mxn": 21500,
//                                         "nok": 11500,
//                                         "nzd": 1900,
//                                         "pln": 4200,
//                                         "sek": 11000
//                                     },
//                                     "is_founding": false,
//                                     "display_name": "Positive Sum (Monthly)",
//                                     "display_price": "$11"
//                                 },
//                                 {
//                                     "sku": "yCskBhBV15sWZT7Ckp",
//                                     "publication_id": "1194762",
//                                     "is_active": true,
//                                     "price_base_units": 11000,
//                                     "currency_alpha3": "usd",
//                                     "period": "year",
//                                     "created_at": "2025-05-29T17:35:50.732Z",
//                                     "updated_at": "2025-05-29T17:35:50.732Z",
//                                     "id": "55248",
//                                     "payout_amount_base_units": 800,
//                                     "alternate_currencies": {
//                                         "aud": 17500,
//                                         "brl": 63000,
//                                         "cad": 15500,
//                                         "chf": 9500,
//                                         "dkk": 73000,
//                                         "eur": 10000,
//                                         "gbp": 8500,
//                                         "mxn": 213500,
//                                         "nok": 112000,
//                                         "nzd": 18500,
//                                         "pln": 41500,
//                                         "sek": 106500
//                                     },
//                                     "is_founding": false,
//                                     "display_name": "Positive Sum (Yearly)",
//                                     "display_price": "$110"
//                                 }
//                             ],
//                             "iap_advanced_plans": [
//                                 {
//                                     "sku": "VajhJQ5OQpwY5tB9Nh",
//                                     "publication_id": "1194762",
//                                     "is_active": true,
//                                     "price_base_units": 1100,
//                                     "currency_alpha3": "usd",
//                                     "period": "month",
//                                     "created_at": "2025-05-29T17:35:50.716Z",
//                                     "updated_at": "2025-05-29T17:35:50.716Z",
//                                     "id": "55247",
//                                     "payout_amount_base_units": 80,
//                                     "alternate_currencies": {
//                                         "aud": 1800,
//                                         "brl": 6500,
//                                         "cad": 1600,
//                                         "chf": 1000,
//                                         "dkk": 7500,
//                                         "eur": 1000,
//                                         "gbp": 900,
//                                         "mxn": 21500,
//                                         "nok": 11500,
//                                         "nzd": 1900,
//                                         "pln": 4200,
//                                         "sek": 11000
//                                     },
//                                     "is_founding": false,
//                                     "display_name": "Positive Sum (Monthly)",
//                                     "display_price": "$11"
//                                 },
//                                 {
//                                     "sku": "yCskBhBV15sWZT7Ckp",
//                                     "publication_id": "1194762",
//                                     "is_active": true,
//                                     "price_base_units": 11000,
//                                     "currency_alpha3": "usd",
//                                     "period": "year",
//                                     "created_at": "2025-05-29T17:35:50.732Z",
//                                     "updated_at": "2025-05-29T17:35:50.732Z",
//                                     "id": "55248",
//                                     "payout_amount_base_units": 800,
//                                     "alternate_currencies": {
//                                         "aud": 17500,
//                                         "brl": 63000,
//                                         "cad": 15500,
//                                         "chf": 9500,
//                                         "dkk": 73000,
//                                         "eur": 10000,
//                                         "gbp": 8500,
//                                         "mxn": 213500,
//                                         "nok": 112000,
//                                         "nzd": 18500,
//                                         "pln": 41500,
//                                         "sek": 106500
//                                     },
//                                     "is_founding": false,
//                                     "display_name": "Positive Sum (Yearly)",
//                                     "display_price": "$110"
//                                 }
//                             ],
//                             "founding_plan_name_english": "Founding Member",
//                             "iap_founding_plan": {
//                                 "name": "Founding Member",
//                                 "minimum_amount": 11000,
//                                 "suggested_amount": 30000,
//                                 "currency_alpha3": "usd",
//                                 "alternate_currencies": {
//                                     "aud": {
//                                         "minimum_amount": 16000,
//                                         "suggested_amount": 43500
//                                     },
//                                     "brl": {
//                                         "minimum_amount": 58000,
//                                         "suggested_amount": 158000
//                                     },
//                                     "cad": {
//                                         "minimum_amount": 15500,
//                                         "suggested_amount": 41500
//                                     },
//                                     "chf": {
//                                         "minimum_amount": 9000,
//                                         "suggested_amount": 23500
//                                     },
//                                     "dkk": {
//                                         "minimum_amount": 70000,
//                                         "suggested_amount": 190000
//                                     },
//                                     "eur": {
//                                         "minimum_amount": 9500,
//                                         "suggested_amount": 25500
//                                     },
//                                     "gbp": {
//                                         "minimum_amount": 8500,
//                                         "suggested_amount": 22500
//                                     },
//                                     "mxn": {
//                                         "minimum_amount": 192000,
//                                         "suggested_amount": 522500
//                                     },
//                                     "nok": {
//                                         "minimum_amount": 107500,
//                                         "suggested_amount": 293500
//                                     },
//                                     "nzd": {
//                                         "minimum_amount": 18500,
//                                         "suggested_amount": 50500
//                                     },
//                                     "pln": {
//                                         "minimum_amount": 39500,
//                                         "suggested_amount": 107500
//                                     },
//                                     "sek": {
//                                         "minimum_amount": 99500,
//                                         "suggested_amount": 271500
//                                     }
//                                 }
//                             },
//                             "draft_plans": [
//                                 {
//                                     "id": "yearly80usd",
//                                     "object": "plan",
//                                     "active": true,
//                                     "aggregate_usage": null,
//                                     "amount": 8000,
//                                     "amount_decimal": "8000",
//                                     "billing_scheme": "per_unit",
//                                     "created": 1748540077,
//                                     "currency": "usd",
//                                     "interval": "year",
//                                     "interval_count": 1,
//                                     "livemode": true,
//                                     "metadata": {
//                                         "substack": "yes"
//                                     },
//                                     "meter": null,
//                                     "nickname": "$80 a year",
//                                     "product": "prod_SOyI6vIH9Lx4ng",
//                                     "tiers": null,
//                                     "tiers_mode": null,
//                                     "transform_usage": null,
//                                     "trial_period_days": null,
//                                     "usage_type": "licensed",
//                                     "currency_options": {
//                                         "aud": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 12000,
//                                             "unit_amount_decimal": "12000"
//                                         },
//                                         "brl": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 44500,
//                                             "unit_amount_decimal": "44500"
//                                         },
//                                         "cad": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 11000,
//                                             "unit_amount_decimal": "11000"
//                                         },
//                                         "chf": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 6500,
//                                             "unit_amount_decimal": "6500"
//                                         },
//                                         "dkk": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 51000,
//                                             "unit_amount_decimal": "51000"
//                                         },
//                                         "eur": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 7000,
//                                             "unit_amount_decimal": "7000"
//                                         },
//                                         "gbp": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 6000,
//                                             "unit_amount_decimal": "6000"
//                                         },
//                                         "mxn": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 144500,
//                                             "unit_amount_decimal": "144500"
//                                         },
//                                         "nok": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 81000,
//                                             "unit_amount_decimal": "81000"
//                                         },
//                                         "nzd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 14000,
//                                             "unit_amount_decimal": "14000"
//                                         },
//                                         "pln": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 29000,
//                                             "unit_amount_decimal": "29000"
//                                         },
//                                         "sek": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 74000,
//                                             "unit_amount_decimal": "74000"
//                                         },
//                                         "usd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 8000,
//                                             "unit_amount_decimal": "8000"
//                                         }
//                                     }
//                                 },
//                                 {
//                                     "id": "monthly8usd",
//                                     "object": "plan",
//                                     "active": true,
//                                     "aggregate_usage": null,
//                                     "amount": 800,
//                                     "amount_decimal": "800",
//                                     "billing_scheme": "per_unit",
//                                     "created": 1748540076,
//                                     "currency": "usd",
//                                     "interval": "month",
//                                     "interval_count": 1,
//                                     "livemode": true,
//                                     "metadata": {
//                                         "substack": "yes"
//                                     },
//                                     "meter": null,
//                                     "nickname": "$8 a month",
//                                     "product": "prod_SOyI5XfMAvvRkI",
//                                     "tiers": null,
//                                     "tiers_mode": null,
//                                     "transform_usage": null,
//                                     "trial_period_days": null,
//                                     "usage_type": "licensed",
//                                     "currency_options": {
//                                         "aud": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 1200,
//                                             "unit_amount_decimal": "1200"
//                                         },
//                                         "brl": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 4500,
//                                             "unit_amount_decimal": "4500"
//                                         },
//                                         "cad": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 1100,
//                                             "unit_amount_decimal": "1100"
//                                         },
//                                         "chf": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 700,
//                                             "unit_amount_decimal": "700"
//                                         },
//                                         "dkk": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 5500,
//                                             "unit_amount_decimal": "5500"
//                                         },
//                                         "eur": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 700,
//                                             "unit_amount_decimal": "700"
//                                         },
//                                         "gbp": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 600,
//                                             "unit_amount_decimal": "600"
//                                         },
//                                         "mxn": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 14500,
//                                             "unit_amount_decimal": "14500"
//                                         },
//                                         "nok": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 8500,
//                                             "unit_amount_decimal": "8500"
//                                         },
//                                         "nzd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 1400,
//                                             "unit_amount_decimal": "1400"
//                                         },
//                                         "pln": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 2900,
//                                             "unit_amount_decimal": "2900"
//                                         },
//                                         "sek": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 7500,
//                                             "unit_amount_decimal": "7500"
//                                         },
//                                         "usd": {
//                                             "custom_unit_amount": null,
//                                             "tax_behavior": "unspecified",
//                                             "unit_amount": 800,
//                                             "unit_amount_decimal": "800"
//                                         }
//                                     }
//                                 },
//                                 {
//                                     "id": "founding30000usd",
//                                     "name": "founding30000usd",
//                                     "nickname": "founding30000usd",
//                                     "active": true,
//                                     "amount": 30000,
//                                     "currency": "usd",
//                                     "interval": "year",
//                                     "interval_count": 1,
//                                     "metadata": {
//                                         "substack": "yes",
//                                         "founding": "yes",
//                                         "no_coupons": "yes",
//                                         "short_description": "Founding Member",
//                                         "short_description_english": "Founding Member",
//                                         "minimum": "8000",
//                                         "minimum_local": {
//                                             "aud": 11500,
//                                             "brl": 42500,
//                                             "cad": 11000,
//                                             "chf": 6500,
//                                             "dkk": 51000,
//                                             "eur": 7000,
//                                             "gbp": 6000,
//                                             "mxn": 139500,
//                                             "nok": 78500,
//                                             "nzd": 13500,
//                                             "pln": 29000,
//                                             "sek": 72500,
//                                             "usd": 8000
//                                         }
//                                     },
//                                     "currency_options": {
//                                         "aud": {
//                                             "unit_amount": 43500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "brl": {
//                                             "unit_amount": 158000,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "cad": {
//                                             "unit_amount": 41500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "chf": {
//                                             "unit_amount": 23500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "dkk": {
//                                             "unit_amount": 190000,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "eur": {
//                                             "unit_amount": 25500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "gbp": {
//                                             "unit_amount": 22500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "mxn": {
//                                             "unit_amount": 522500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "nok": {
//                                             "unit_amount": 293500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "nzd": {
//                                             "unit_amount": 50500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "pln": {
//                                             "unit_amount": 107500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "sek": {
//                                             "unit_amount": 271500,
//                                             "tax_behavior": "unspecified"
//                                         },
//                                         "usd": {
//                                             "unit_amount": 30000,
//                                             "tax_behavior": "unspecified"
//                                         }
//                                     }
//                                 }
//                             ],
//                             "bundles": [],
//                             "base_url": "https://abio.substack.com",
//                             "hostname": "abio.substack.com",
//                             "is_on_substack": false,
//                             "spotify_podcast_settings": null,
//                             "podcastPalette": {
//                                 "DarkMuted": {
//                                     "population": 72,
//                                     "rgb": [
//                                         73,
//                                         153,
//                                         137
//                                     ]
//                                 },
//                                 "DarkVibrant": {
//                                     "population": 6013,
//                                     "rgb": [
//                                         4,
//                                         100,
//                                         84
//                                     ]
//                                 },
//                                 "LightMuted": {
//                                     "population": 7,
//                                     "rgb": [
//                                         142,
//                                         198,
//                                         186
//                                     ]
//                                 },
//                                 "LightVibrant": {
//                                     "population": 3,
//                                     "rgb": [
//                                         166,
//                                         214,
//                                         206
//                                     ]
//                                 },
//                                 "Muted": {
//                                     "population": 6,
//                                     "rgb": [
//                                         92,
//                                         164,
//                                         156
//                                     ]
//                                 },
//                                 "Vibrant": {
//                                     "population": 5,
//                                     "rgb": [
//                                         76,
//                                         164,
//                                         146
//                                     ]
//                                 }
//                             },
//                             "pageThemes": {
//                                 "podcast": null
//                             },
//                             "multiple_pins": true,
//                             "live_subscriber_counts": false,
//                             "supports_ip_content_unlock": false,
//                             "appTheme": {
//                                 "colors": {
//                                     "accent": {
//                                         "name": "#67bdfc",
//                                         "primary": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 1
//                                         },
//                                         "primary_hover": {
//                                             "r": 81,
//                                             "g": 170,
//                                             "b": 232,
//                                             "a": 1
//                                         },
//                                         "primary_elevated": {
//                                             "r": 81,
//                                             "g": 170,
//                                             "b": 232,
//                                             "a": 1
//                                         },
//                                         "secondary": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 0.2
//                                         },
//                                         "contrast": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "bg": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 0.2
//                                         },
//                                         "bg_hover": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 0.3
//                                         },
//                                         "dark": {
//                                             "primary": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 1
//                                             },
//                                             "primary_hover": {
//                                                 "r": 129,
//                                                 "g": 208,
//                                                 "b": 255,
//                                                 "a": 1
//                                             },
//                                             "primary_elevated": {
//                                                 "r": 129,
//                                                 "g": 208,
//                                                 "b": 255,
//                                                 "a": 1
//                                             },
//                                             "secondary": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 0.2
//                                             },
//                                             "contrast": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 1
//                                             },
//                                             "bg": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 0.2
//                                             },
//                                             "bg_hover": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 0.3
//                                             }
//                                         }
//                                     },
//                                     "fg": {
//                                         "primary": {
//                                             "r": 0,
//                                             "g": 0,
//                                             "b": 0,
//                                             "a": 0.8
//                                         },
//                                         "secondary": {
//                                             "r": 0,
//                                             "g": 0,
//                                             "b": 0,
//                                             "a": 0.6
//                                         },
//                                         "tertiary": {
//                                             "r": 0,
//                                             "g": 0,
//                                             "b": 0,
//                                             "a": 0.4
//                                         },
//                                         "accent": {
//                                             "r": 0,
//                                             "g": 125,
//                                             "b": 184,
//                                             "a": 1
//                                         },
//                                         "dark": {
//                                             "primary": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 0.9
//                                             },
//                                             "secondary": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 0.6
//                                             },
//                                             "tertiary": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 0.4
//                                             },
//                                             "accent": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 1
//                                             }
//                                         }
//                                     },
//                                     "bg": {
//                                         "name": "#ffffff",
//                                         "hue": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 0
//                                         },
//                                         "tint": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 0
//                                         },
//                                         "primary": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "primary_hover": {
//                                             "r": 250,
//                                             "g": 250,
//                                             "b": 250,
//                                             "a": 1
//                                         },
//                                         "primary_elevated": {
//                                             "r": 250,
//                                             "g": 250,
//                                             "b": 250,
//                                             "a": 1
//                                         },
//                                         "secondary": {
//                                             "r": 238,
//                                             "g": 238,
//                                             "b": 238,
//                                             "a": 1
//                                         },
//                                         "secondary_elevated": {
//                                             "r": 206.90096477355226,
//                                             "g": 206.90096477355175,
//                                             "b": 206.9009647735519,
//                                             "a": 1
//                                         },
//                                         "tertiary": {
//                                             "r": 219,
//                                             "g": 219,
//                                             "b": 219,
//                                             "a": 1
//                                         },
//                                         "quaternary": {
//                                             "r": 182,
//                                             "g": 182,
//                                             "b": 182,
//                                             "a": 1
//                                         },
//                                         "dark": {
//                                             "primary": {
//                                                 "r": 22,
//                                                 "g": 23,
//                                                 "b": 24,
//                                                 "a": 1
//                                             },
//                                             "primary_hover": {
//                                                 "r": 27,
//                                                 "g": 28,
//                                                 "b": 29,
//                                                 "a": 1
//                                             },
//                                             "primary_elevated": {
//                                                 "r": 27,
//                                                 "g": 28,
//                                                 "b": 29,
//                                                 "a": 1
//                                             },
//                                             "secondary": {
//                                                 "r": 35,
//                                                 "g": 37,
//                                                 "b": 37,
//                                                 "a": 1
//                                             },
//                                             "secondary_elevated": {
//                                                 "r": 41.35899397549579,
//                                                 "g": 43.405356429195315,
//                                                 "b": 43.40489285041963,
//                                                 "a": 1
//                                             },
//                                             "tertiary": {
//                                                 "r": 54,
//                                                 "g": 55,
//                                                 "b": 55,
//                                                 "a": 1
//                                             },
//                                             "quaternary": {
//                                                 "r": 90,
//                                                 "g": 91,
//                                                 "b": 91,
//                                                 "a": 1
//                                             }
//                                         }
//                                     }
//                                 }
//                             },
//                             "portalAppTheme": {
//                                 "colors": {
//                                     "accent": {
//                                         "name": "#67BDFC",
//                                         "primary": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 1
//                                         },
//                                         "primary_hover": {
//                                             "r": 78,
//                                             "g": 178,
//                                             "b": 252,
//                                             "a": 1
//                                         },
//                                         "primary_elevated": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 1
//                                         },
//                                         "secondary": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 1
//                                         },
//                                         "contrast": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "bg": {
//                                             "r": 255,
//                                             "g": 103,
//                                             "b": 25,
//                                             "a": 0.2
//                                         },
//                                         "bg_hover": {
//                                             "r": 255,
//                                             "g": 103,
//                                             "b": 25,
//                                             "a": 0.3
//                                         },
//                                         "dark": {
//                                             "primary": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 1
//                                             },
//                                             "primary_hover": {
//                                                 "r": 129,
//                                                 "g": 208,
//                                                 "b": 255,
//                                                 "a": 1
//                                             },
//                                             "primary_elevated": {
//                                                 "r": 129,
//                                                 "g": 208,
//                                                 "b": 255,
//                                                 "a": 1
//                                             },
//                                             "secondary": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 0.2
//                                             },
//                                             "contrast": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 1
//                                             },
//                                             "bg": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 0.2
//                                             },
//                                             "bg_hover": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 0.3
//                                             }
//                                         }
//                                     },
//                                     "fg": {
//                                         "primary": {
//                                             "r": 54,
//                                             "g": 55,
//                                             "b": 55,
//                                             "a": 1
//                                         },
//                                         "secondary": {
//                                             "r": 134,
//                                             "g": 135,
//                                             "b": 135,
//                                             "a": 1
//                                         },
//                                         "tertiary": {
//                                             "r": 146,
//                                             "g": 146,
//                                             "b": 146,
//                                             "a": 1
//                                         },
//                                         "accent": {
//                                             "r": 103,
//                                             "g": 189,
//                                             "b": 252,
//                                             "a": 1
//                                         },
//                                         "dark": {
//                                             "primary": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 0.9
//                                             },
//                                             "secondary": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 0.6
//                                             },
//                                             "tertiary": {
//                                                 "r": 255,
//                                                 "g": 255,
//                                                 "b": 255,
//                                                 "a": 0.4
//                                             },
//                                             "accent": {
//                                                 "r": 103,
//                                                 "g": 189,
//                                                 "b": 252,
//                                                 "a": 1
//                                             }
//                                         }
//                                     },
//                                     "bg": {
//                                         "name": "#ffffff",
//                                         "hue": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "tint": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "primary": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "primary_hover": {
//                                             "r": 240,
//                                             "g": 240,
//                                             "b": 240,
//                                             "a": 1
//                                         },
//                                         "primary_elevated": {
//                                             "r": 255,
//                                             "g": 255,
//                                             "b": 255,
//                                             "a": 1
//                                         },
//                                         "secondary": {
//                                             "r": 240,
//                                             "g": 240,
//                                             "b": 240,
//                                             "a": 1
//                                         },
//                                         "secondary_elevated": {
//                                             "r": 240,
//                                             "g": 240,
//                                             "b": 240,
//                                             "a": 1
//                                         },
//                                         "tertiary": {
//                                             "r": 221,
//                                             "g": 221,
//                                             "b": 221,
//                                             "a": 1
//                                         },
//                                         "quaternary": {
//                                             "r": 183,
//                                             "g": 183,
//                                             "b": 183,
//                                             "a": 1
//                                         },
//                                         "dark": {
//                                             "primary": {
//                                                 "r": 22,
//                                                 "g": 23,
//                                                 "b": 24,
//                                                 "a": 1
//                                             },
//                                             "primary_hover": {
//                                                 "r": 27,
//                                                 "g": 28,
//                                                 "b": 29,
//                                                 "a": 1
//                                             },
//                                             "primary_elevated": {
//                                                 "r": 27,
//                                                 "g": 28,
//                                                 "b": 29,
//                                                 "a": 1
//                                             },
//                                             "secondary": {
//                                                 "r": 35,
//                                                 "g": 37,
//                                                 "b": 37,
//                                                 "a": 1
//                                             },
//                                             "secondary_elevated": {
//                                                 "r": 41.35899397549579,
//                                                 "g": 43.405356429195315,
//                                                 "b": 43.40489285041963,
//                                                 "a": 1
//                                             },
//                                             "tertiary": {
//                                                 "r": 54,
//                                                 "g": 55,
//                                                 "b": 55,
//                                                 "a": 1
//                                             },
//                                             "quaternary": {
//                                                 "r": 90,
//                                                 "g": 91,
//                                                 "b": 91,
//                                                 "a": 1
//                                             }
//                                         }
//                                     }
//                                 }
//                             }
//                         },
//                         "post": {
//                             "id": 183806048,
//                             "publication_id": 1194762,
//                             "title": "How Working-Class People Talk About Uber Is Not How the Media Covers It",
//                             "social_title": null,
//                             "search_engine_title": null,
//                             "search_engine_description": null,
//                             "type": "newsletter",
//                             "slug": "how-working-class-people-talk-about",
//                             "post_date": "2026-01-07T19:05:10.610Z",
//                             "audience": "everyone",
//                             "podcast_duration": null,
//                             "video_upload_id": null,
//                             "write_comment_permissions": "everyone",
//                             "should_send_free_preview": false,
//                             "free_unlock_required": false,
//                             "default_comment_sort": null,
//                             "canonical_url": "https://abio.substack.com/p/how-working-class-people-talk-about",
//                             "section_id": null,
//                             "podcast_art_url": null,
//                             "is_published": true,
//                             "live_stream_id": null,
//                             "restacks": 26,
//                             "top_exclusions": [],
//                             "pins": [],
//                             "is_section_pinned": false,
//                             "has_shareable_clips": false,
//                             "section_slug": null,
//                             "section_name": null,
//                             "reactions": {
//                                 "❤": 260
//                             },
//                             "subtitle": "The people writing about whether gig work is exploitative have usually never competed for a restaurant shift. The people most helped by a new tech are often the least represented in policy debates.",
//                             "cover_image": "https://images.unsplash.com/photo-1600320254374-ce2d293c324e?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wzMDAzMzh8MHwxfHNlYXJjaHwzfHx1YmVyfGVufDB8fHx8MTc2NzgwODI0N3ww&ixlib=rb-4.1.0&q=80&w=1080",
//                             "cover_image_is_square": false,
//                             "cover_image_is_explicit": false,
//                             "podcast_url": null,
//                             "videoUpload": null,
//                             "podcastFields": {
//                                 "post_id": 183806048,
//                                 "podcast_episode_number": null,
//                                 "podcast_season_number": null,
//                                 "podcast_episode_type": null,
//                                 "should_syndicate_to_other_feed": null,
//                                 "syndicate_to_section_id": null,
//                                 "hide_from_feed": false,
//                                 "free_podcast_url": null,
//                                 "free_podcast_duration": null
//                             },
//                             "podcast_upload_id": null,
//                             "podcast_preview_upload_id": null,
//                             "podcastUpload": null,
//                             "podcastPreviewUpload": null,
//                             "voiceover_upload_id": null,
//                             "voiceoverUpload": null,
//                             "has_voiceover": false,
//                             "description": "The people writing about whether gig work is exploitative have usually never competed for a restaurant shift. The people most helped by a new tech are often the least represented in policy debates.",
//                             "body_json": null,
//                             "body_html": null,
//                             "truncated_body_text": "I waited tables for six years at $2.13 an hour. Now I'm a policy researcher in DC. These two worlds talk about technology completely differently. Let me use Uber driving as an example.",
//                             "wordcount": 642,
//                             "postTags": [
//                                 {
//                                     "id": "0a007c22-f608-4dc6-a952-8fa27dd31915",
//                                     "publication_id": 1194762,
//                                     "name": "news",
//                                     "slug": "news",
//                                     "hidden": false
//                                 },
//                                 {
//                                     "id": "24d58f03-64f9-432b-9086-95867a6fa33b",
//                                     "publication_id": 1194762,
//                                     "name": "pessimism",
//                                     "slug": "pessimism",
//                                     "hidden": false
//                                 },
//                                 {
//                                     "id": "25906212-58a6-4172-91cc-66eb24adc57c",
//                                     "publication_id": 1194762,
//                                     "name": "progress",
//                                     "slug": "progress",
//                                     "hidden": false
//                                 },
//                                 {
//                                     "id": "36ac7cc2-9168-4646-a269-cfc1bca59cf4",
//                                     "publication_id": 1194762,
//                                     "name": "society",
//                                     "slug": "society",
//                                     "hidden": false
//                                 },
//                                 {
//                                     "id": "38871425-2a94-4f9e-b057-d901e70880b2",
//                                     "publication_id": 1194762,
//                                     "name": "abundance",
//                                     "slug": "abundance",
//                                     "hidden": false
//                                 },
//                                 {
//                                     "id": "628e152c-ba3f-4e0d-869f-816a209364c5",
//                                     "publication_id": 1194762,
//                                     "name": "tech",
//                                     "slug": "tech",
//                                     "hidden": false
//                                 }
//                             ],
//                             "teaser_post_eligible": true,
//                             "postCountryBlocks": [],
//                             "headlineTest": null,
//                             "coverImagePalette": null,
//                             "publishedBylines": [
//                                 {
//                                     "id": 349629,
//                                     "name": "Abi Olvera",
//                                     "handle": "abio",
//                                     "previous_name": "Abigail Olvera",
//                                     "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/550e023c-2e8e-440f-91e8-6d32872d8d5f_1123x1125.png",
//                                     "bio": "Optimistic takes on why things cost so much, why institutions make mistakes, and what AI is actually changing. Ex-diplomat. I lead DC Abundance. Golden Gate Institute Research Director. Emergent Ventures grantee. Views my own.",
//                                     "profile_set_up_at": "2022-09-16T15:14:52.491Z",
//                                     "reader_installed_at": "2022-09-16T15:13:28.144Z",
//                                     "publicationUsers": [
//                                         {
//                                             "id": 1148952,
//                                             "user_id": 349629,
//                                             "publication_id": 1194762,
//                                             "role": "admin",
//                                             "public": true,
//                                             "is_primary": true,
//                                             "publication": {
//                                                 "id": 1194762,
//                                                 "name": "Positive Sum",
//                                                 "subdomain": "abio",
//                                                 "custom_domain": null,
//                                                 "custom_domain_optional": false,
//                                                 "hero_text": "Optimistic takes on why things cost so much, why institutions make mistakes, and what AI is actually changing.",
//                                                 "logo_url": null,
//                                                 "author_id": 349629,
//                                                 "primary_user_id": 349629,
//                                                 "theme_var_background_pop": "#67BDFC",
//                                                 "created_at": "2022-11-16T02:38:49.439Z",
//                                                 "email_from_name": "Abi Olvera's Newsletter",
//                                                 "copyright": "Abigail Olvera",
//                                                 "founding_plan_name": "Founding Member",
//                                                 "community_enabled": true,
//                                                 "invite_only": false,
//                                                 "payments_state": "enabled",
//                                                 "language": null,
//                                                 "explicit": false,
//                                                 "homepage_type": "magaziney",
//                                                 "is_personal_mode": false
//                                             }
//                                         }
//                                     ],
//                                     "twitter_screen_name": "Abi0lvera",
//                                     "is_guest": false,
//                                     "bestseller_tier": null,
//                                     "status": {
//                                         "bestsellerTier": null,
//                                         "subscriberTier": 1,
//                                         "leaderboard": null,
//                                         "vip": false,
//                                         "badge": {
//                                             "type": "subscriber",
//                                             "tier": 1,
//                                             "accent_colors": null
//                                         },
//                                         "paidPublicationIds": [
//                                             2252,
//                                             1793203,
//                                             159185
//                                         ],
//                                         "subscriber": null
//                                     },
//                                     "primary_publication": {
//                                         "id": 1194762,
//                                         "subdomain": "abio",
//                                         "custom_domain_optional": false,
//                                         "name": "Positive Sum",
//                                         "author_id": 349629,
//                                         "user_id": 349629,
//                                         "handles_enabled": false,
//                                         "explicit": false,
//                                         "is_personal_mode": false,
//                                         "payments_state": "enabled",
//                                         "pledges_enabled": false
//                                     }
//                                 }
//                             ],
//                             "reaction": false,
//                             "reaction_count": 260,
//                             "comment_count": 58,
//                             "child_comment_count": 17,
//                             "audio_items": [
//                                 {
//                                     "post_id": 183806048,
//                                     "voice_id": "en-US-OnyxTurboMultilingualNeural",
//                                     "audio_url": "https://substack-video.s3.amazonaws.com/video_upload/post/183806048/tts/ed6da285-093d-460f-9486-1ad809064428/en-US-OnyxTurboMultilingualNeural.mp3",
//                                     "type": "tts",
//                                     "status": "completed"
//                                 }
//                             ],
//                             "is_geoblocked": false,
//                             "hasCashtag": false,
//                             "is_saved": false,
//                             "saved_at": null,
//                             "is_viewed": false,
//                             "read_progress": 0,
//                             "max_read_progress": 0,
//                             "audio_progress": 0,
//                             "max_audio_progress": 0,
//                             "video_progress": 0,
//                             "max_video_progress": 0,
//                             "restacked": false
//                         },
//                         "postSelection": null,
//                         "postSelectionTheme": null,
//                         "postImageSelection": null,
//                         "clipInfo": null,
//                         "mediaClip": null
//                     }
//                 ],
//                 "user_bestseller_tier": null,
//                 "userStatus": {
//                     "bestsellerTier": null,
//                     "subscriberTier": 1,
//                     "leaderboard": null,
//                     "vip": false,
//                     "badge": {
//                         "type": "subscriber",
//                         "tier": 1,
//                         "accent_colors": null
//                     },
//                     "paidPublicationIds": [
//                         2252,
//                         1793203,
//                         159185
//                     ],
//                     "subscriber": null
//                 },
//                 "user_primary_publication": {
//                     "id": 1194762,
//                     "subdomain": "abio",
//                     "custom_domain_optional": false,
//                     "name": "Positive Sum",
//                     "author_id": 349629,
//                     "user_id": 349629,
//                     "handles_enabled": false,
//                     "explicit": false,
//                     "is_personal_mode": false,
//                     "payments_state": "enabled",
//                     "pledges_enabled": false
//                 }
//             },
//             "parentComments": [],
//             "canReply": true,
//             "isMuted": false,
//             "trackingParameters": {
//                 "item_primary_entity_key": "c-196576534",
//                 "item_entity_key": "c-196576534",
//                 "item_type": "comment",
//                 "item_comment_id": 196576534,
//                 "item_content_user_id": 349629,
//                 "item_content_timestamp": "2026-01-07T19:19:50.470Z",
//                 "item_context_type": "comment_like",
//                 "item_context_type_bucket": "notes",
//                 "item_context_timestamp": "2026-01-07T19:19:50.470Z",
//                 "item_context_user_id": 10472909,
//                 "item_context_user_ids": [
//                     10472909
//                 ],
//                 "item_can_reply": true,
//                 "item_is_fresh": true,
//                 "item_last_impression_at": null,
//                 "item_source": "model",
//                 "item_model_rank": 1,
//                 "item_model_score": 0.5,
//                 "source_comment_id": 209255871,
//                 "item_page": 0,
//                 "item_page_rank": 1,
//                 "tab_id": "for-you",
//                 "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//                 "impression_id": "dc8f3b9c-1365-46ac-9f27-94a6f5684421",
//                 "followed_user_count": 224,
//                 "subscribed_publication_count": 205,
//                 "is_following": false,
//                 "is_explicitly_subscribed": false,
//                 "note_velocity_factor": 1.078759961279,
//                 "note_delay_seconds": 141,
//                 "note_notes_per_hour": 2819.426967,
//                 "item_current_reaction_count": 659,
//                 "item_current_restack_count": 34,
//                 "item_current_reply_count": 18
//             },
//             "canShowFollowUpsell": false
//         },
//         {
//             "entity_key": "c-208836424",
//             "type": "comment",
//             "context": {
//                 "type": "comment_like",
//                 "typeBucket": "notes",
//                 "timestamp": "2026-02-02T15:14:11.978Z",
//                 "users": [
//                     {
//                         "id": 14807526,
//                         "name": "Gary Marcus",
//                         "handle": "garymarcus",
//                         "previous_name": null,
//                         "photo_url": "https://substackcdn.com/image/fetch/$s_!Ka51!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F8fb2e48c-be2a-4db7-b68c-90300f00fd1e_1668x1456.jpeg",
//                         "bio": "Scientist, author and entrepreneur, known as a leading voice in AI. Six books including The Algebraic Mind, Rebooting AI, and Taming Silicon Valley; NYU Professor Emeritus.",
//                         "profile_set_up_at": "2022-05-14T14:01:17.198Z",
//                         "reader_installed_at": "2022-05-14T13:59:03.190Z",
//                         "bestseller_tier": 1000,
//                         "status": {
//                             "bestsellerTier": 1000,
//                             "subscriberTier": null,
//                             "leaderboard": {
//                                 "ranking": "trending",
//                                 "rank": 6,
//                                 "publicationName": "Marcus on AI",
//                                 "label": "Technology",
//                                 "categoryId": "4",
//                                 "publicationId": 888615
//                             },
//                             "vip": false,
//                             "badge": {
//                                 "type": "bestseller",
//                                 "tier": 1000
//                             },
//                             "paidPublicationIds": [],
//                             "subscriber": null
//                         },
//                         "primary_publication": {
//                             "id": 888615,
//                             "subdomain": "garymarcus",
//                             "custom_domain_optional": false,
//                             "name": "Marcus on AI",
//                             "author_id": 14807526,
//                             "user_id": 14807526,
//                             "handles_enabled": false,
//                             "explicit": false,
//                             "is_personal_mode": false,
//                             "payments_state": "enabled",
//                             "pledges_enabled": false
//                         }
//                     }
//                 ],
//                 "fallbackReason": "//TODO: create fallback reasons",
//                 "fallbackUrl": null,
//                 "isFresh": true,
//                 "source": "model",
//                 "model_rank": 2,
//                 "model_score": 0.3333333333333333,
//                 "page": 0,
//                 "page_rank": 2
//             },
//             "publication": null,
//             "post": null,
//             "comment": {
//                 "name": "Arvind Narayanan",
//                 "handle": "aisnakeoil",
//                 "photo_url": "https://substackcdn.com/image/fetch/$s_!bVLI!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fdd0d6558-256e-46c4-b2c5-7cf7f808a9c9_693x693.jpeg",
//                 "id": 208836424,
//                 "body": "Why do coding agents work so well and what would it take to replicate their success in other domains? One important and under-appreciated reason is that agentic coding is a type of neurosymbolic AI.\n\n\n\nThe main weakness of LLMs is that they are statistical machines and struggle at tasks involving long chains of logic / symbol manipulation. Of course, traditional code is the opposite. The magic of agentic coding is that it fuses the two — there is a lot of code execution during code generation. This is a subtle point so let me spell it out.\n\n* Most obviously, agents run the generated code itself, run tests, etc. This makes coding a verifiable domain. It is well known that in verifiable domains, inference scaling is highly effective as agents can fix their own mistakes. It also allows reinforcement learning to be highly effective.\n\n* Next, code generation often takes advantage of existing symbolic tools like compilers that have been optimized and perfected over decades. Imagine if LLMs had to directly output binary code instead. (They sometimes can, and it's a cool trick, but it's no way to do software engineering.)\n\n* IMO the biggest neurosymbolic unlock is the shell, which allows a dramatic expansion in capabilities by using existing tools to effectively do complex editing tasks. Many of us remember the feeling of wizardry when we gained shell fluency. LLMs are able to pick up shell knowledge and best practices through pre-training because it is extensively documented on places like StackOverflow.\n\n* Finally, more complex agentic coding tasks often involve LLMs writing code that in turn invokes LLMs. In principle you can have an arbitrary depth of recursion between statistical and symbolic systems.\n\nNeurosymbolic AI is a touchy topic and many people have their own favored conception of what it should look like. And admittedly agentic coding uses really crude patterns, with LLMs and code being loosely coupled. But the point is — it works! LLMs are able to use the giant warehouse of tools that humans have built over the decades to reach ever-increasing levels of abstraction and complexity.\n\nTo build agentic systems in other domains, here’s what we need. First, it must be a verifiable domain. Math is and writing isn’t. There’s no getting around that. Provided we’re in a friendly domain, it all comes down to whether we can build a symbolic toolbox, and how well LLMs can be trained to use that toolbox. IMO this is where the alpha will be, more so than in LLM capabilities themselves.",
//                 "body_json": {
//                     "type": "doc",
//                     "attrs": {
//                         "schemaVersion": "v1"
//                     },
//                     "content": [
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Why do coding agents work so well and what would it take to replicate their success in other domains? One important and under-appreciated reason is that agentic coding is a type of neurosymbolic AI."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "The main weakness of LLMs is that they are statistical machines and struggle at tasks involving long chains of logic / symbol manipulation. Of course, traditional code is the opposite. The magic of agentic coding is that it fuses the two — there is a lot of code "
//                                 },
//                                 {
//                                     "type": "text",
//                                     "marks": [
//                                         {
//                                             "type": "italic"
//                                         }
//                                     ],
//                                     "text": "execution"
//                                 },
//                                 {
//                                     "type": "text",
//                                     "text": " during code generation. This is a subtle point so let me spell it out."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "* Most obviously, agents run the generated code itself, run tests, etc. This makes coding a verifiable domain. It is well known that in verifiable domains, inference scaling is highly effective as agents can fix their own mistakes. It also allows reinforcement learning to be highly effective."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "* Next, code generation often takes advantage of existing symbolic tools like compilers that have been optimized and perfected over decades. Imagine if LLMs had to directly output binary code instead. (They sometimes can, and it's a cool trick, but it's no way to do software engineering.)"
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "* IMO the biggest neurosymbolic unlock is the shell, which allows a dramatic expansion in capabilities by using existing tools to effectively do complex editing tasks. Many of us remember the feeling of wizardry when we gained shell fluency. LLMs are able to pick up shell knowledge and best practices through pre-training because it is extensively documented on places like StackOverflow."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "* Finally, more complex agentic coding tasks often involve LLMs writing code that in turn invokes LLMs. In principle you can have an arbitrary depth of recursion between statistical and symbolic systems."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Neurosymbolic AI is a touchy topic and many people have their own favored conception of what it should look like. And admittedly agentic coding uses really crude patterns, with LLMs and code being loosely coupled. But the point is — it works! LLMs are able to use the giant warehouse of tools that humans have built over the decades to reach ever-increasing levels of abstraction and complexity."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "To build agentic systems in other domains, here’s what we need. First, it must be a verifiable domain. Math is and writing isn’t. There’s no getting around that. Provided we’re in a friendly domain, it all comes down to whether we can build a symbolic toolbox, and how well LLMs can be trained to use that toolbox. IMO this is where the alpha will be, more so than in LLM capabilities themselves."
//                                 }
//                             ]
//                         }
//                     ]
//                 },
//                 "publication_id": null,
//                 "post_id": null,
//                 "user_id": 19265788,
//                 "type": "feed",
//                 "date": "2026-02-02T15:14:11.978Z",
//                 "edited_at": null,
//                 "ancestor_path": "",
//                 "reply_minimum_role": "everyone",
//                 "media_clip_id": null,
//                 "reaction_count": 91,
//                 "reactions": {
//                     "❤": 91
//                 },
//                 "restacks": 6,
//                 "restacked": false,
//                 "children_count": 9,
//                 "attachments": [],
//                 "user_bestseller_tier": null,
//                 "userStatus": {
//                     "bestsellerTier": null,
//                     "subscriberTier": 1,
//                     "leaderboard": null,
//                     "vip": false,
//                     "badge": {
//                         "type": "subscriber",
//                         "tier": 1,
//                         "accent_colors": null
//                     },
//                     "paidPublicationIds": [
//                         883883,
//                         35345,
//                         5247799,
//                         2203516
//                     ],
//                     "subscriber": null
//                 },
//                 "user_primary_publication": {
//                     "id": 1008003,
//                     "subdomain": "aisnakeoil",
//                     "custom_domain": "www.normaltech.ai",
//                     "custom_domain_optional": false,
//                     "name": "AI as Normal Technology",
//                     "logo_url": "https://substack-post-media.s3.amazonaws.com/public/images/6780dda8-5879-4789-bf02-b3ea43a9f85e_1000x1000.png",
//                     "author_id": 891603,
//                     "user_id": 19265788,
//                     "handles_enabled": false,
//                     "explicit": false,
//                     "is_personal_mode": false,
//                     "payments_state": "disabled",
//                     "pledges_enabled": false
//                 }
//             },
//             "parentComments": [],
//             "canReply": true,
//             "isMuted": false,
//             "trackingParameters": {
//                 "item_primary_entity_key": "c-208836424",
//                 "item_entity_key": "c-208836424",
//                 "item_type": "comment",
//                 "item_comment_id": 208836424,
//                 "item_content_user_id": 19265788,
//                 "item_content_timestamp": "2026-02-02T15:14:11.978Z",
//                 "item_context_type": "comment_like",
//                 "item_context_type_bucket": "notes",
//                 "item_context_timestamp": "2026-02-02T15:14:11.978Z",
//                 "item_context_user_id": 14807526,
//                 "item_context_user_ids": [
//                     14807526
//                 ],
//                 "item_can_reply": true,
//                 "item_is_fresh": true,
//                 "item_last_impression_at": null,
//                 "item_source": "model",
//                 "item_model_rank": 2,
//                 "item_model_score": 0.3333333333333333,
//                 "source_comment_id": 209255871,
//                 "item_page": 0,
//                 "item_page_rank": 2,
//                 "tab_id": "for-you",
//                 "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//                 "impression_id": "67f6d86b-404f-4ded-b4ab-0967126fe8eb",
//                 "followed_user_count": 224,
//                 "subscribed_publication_count": 205,
//                 "is_following": true,
//                 "is_explicitly_subscribed": true,
//                 "note_velocity_factor": 1.078759961279,
//                 "note_delay_seconds": 141,
//                 "note_notes_per_hour": 2819.426967,
//                 "item_current_reaction_count": 91,
//                 "item_current_restack_count": 6,
//                 "item_current_reply_count": 9
//             },
//             "canShowFollowUpsell": false
//         },
//         {
//             "entity_key": "c-203269782",
//             "type": "comment",
//             "context": {
//                 "type": "comment_restack",
//                 "typeBucket": "notes",
//                 "timestamp": "2026-01-22T04:02:52.199Z",
//                 "users": [
//                     {
//                         "id": 86,
//                         "name": "Bill Bishop",
//                         "handle": "sinocism",
//                         "previous_name": "Tashi B",
//                         "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/d029e11f-c157-48e1-a8ef-9dfb53746b8d_850x850.png",
//                         "bio": "Writes The Sinocism Newsletter https://sinocism.com about China, tweets @niubi Human of Tashi the Golden Doodle https://tashigd.substack.com Investor in Substack ",
//                         "profile_set_up_at": "2021-04-23T11:36:35.041Z",
//                         "reader_installed_at": "2022-08-05T22:54:07.489Z",
//                         "bestseller_tier": 1000,
//                         "status": {
//                             "bestsellerTier": 1000,
//                             "subscriberTier": 10,
//                             "leaderboard": {
//                                 "ranking": "paid",
//                                 "rank": 5,
//                                 "publicationName": "Sinocism",
//                                 "label": "News",
//                                 "categoryId": "103",
//                                 "publicationId": 2
//                             },
//                             "vip": false,
//                             "badge": {
//                                 "type": "bestseller",
//                                 "tier": 1000
//                             },
//                             "paidPublicationIds": [
//                                 264786,
//                                 2960805,
//                                 6819723,
//                                 2252,
//                                 5966616,
//                                 474083,
//                                 3996179,
//                                 490119,
//                                 22679,
//                                 6667,
//                                 247191,
//                                 471923,
//                                 4790652,
//                                 354253,
//                                 11020,
//                                 321695,
//                                 1495434,
//                                 324097,
//                                 268333,
//                                 1258230,
//                                 1223483,
//                                 1198116,
//                                 769883,
//                                 1518861,
//                                 748248,
//                                 1301210,
//                                 411027,
//                                 1198404,
//                                 793713,
//                                 1093938,
//                                 1857301,
//                                 1989755,
//                                 1254848,
//                                 4106,
//                                 1981201,
//                                 1655750,
//                                 3492025,
//                                 661423,
//                                 2100547,
//                                 1218238,
//                                 2590773,
//                                 2794840,
//                                 69345,
//                                 6369925,
//                                 3011576,
//                                 1083330,
//                                 2079959,
//                                 15795,
//                                 3009844,
//                                 3525780,
//                                 2244049,
//                                 1174934,
//                                 1005023,
//                                 3072903,
//                                 1534895,
//                                 265424,
//                                 3225541,
//                                 2880588,
//                                 3823437,
//                                 247761,
//                                 5035003,
//                                 1487717,
//                                 5931581,
//                                 4424676,
//                                 4533279,
//                                 2660356,
//                                 1851725,
//                                 620446,
//                                 5291885,
//                                 38441,
//                                 5046759,
//                                 2695528,
//                                 10025,
//                                 7677,
//                                 618944,
//                                 2495006,
//                                 1579393,
//                                 5091299,
//                                 791191,
//                                 484474,
//                                 1074587,
//                                 2029781,
//                                 46963,
//                                 500492,
//                                 712667,
//                                 27199,
//                                 1335949,
//                                 2069548,
//                                 2681713,
//                                 4222056,
//                                 37758,
//                                 4385786,
//                                 1899900,
//                                 895339,
//                                 2273840,
//                                 800777,
//                                 1199539,
//                                 41573,
//                                 1353279,
//                                 2337656,
//                                 82416,
//                                 1444443,
//                                 4220,
//                                 192845,
//                                 25216,
//                                 81003,
//                                 1885086,
//                                 48208,
//                                 2585949,
//                                 391912,
//                                 2666251,
//                                 4272766,
//                                 6676577,
//                                 28991,
//                                 4559901,
//                                 1009222,
//                                 2694757,
//                                 244069,
//                                 48206,
//                                 1455037,
//                                 2412034,
//                                 101810,
//                                 3113246,
//                                 5441321,
//                                 868206,
//                                 1010841,
//                                 6771668,
//                                 501423,
//                                 3138298,
//                                 157712,
//                                 835088,
//                                 87281,
//                                 3188037,
//                                 4363663,
//                                 351976,
//                                 6027,
//                                 50481,
//                                 3410682,
//                                 6803067,
//                                 3394151,
//                                 1284013,
//                                 98827,
//                                 1744858,
//                                 976016,
//                                 260347,
//                                 277517,
//                                 1180066,
//                                 1323936,
//                                 1078491,
//                                 3775894,
//                                 3930677,
//                                 6256426,
//                                 3722532,
//                                 527096
//                             ],
//                             "subscriber": null
//                         },
//                         "primary_publication": {
//                             "id": 2,
//                             "subdomain": "sinocism",
//                             "custom_domain": "sinocism.com",
//                             "custom_domain_optional": false,
//                             "name": "Sinocism",
//                             "logo_url": "https://bucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com/public/images/031353ec-20cb-462c-8860-bbd04365b90c_256x256",
//                             "author_id": 86,
//                             "user_id": 86,
//                             "handles_enabled": false,
//                             "explicit": false,
//                             "is_personal_mode": false,
//                             "payments_state": "enabled",
//                             "pledges_enabled": false
//                         }
//                     }
//                 ],
//                 "fallbackReason": "//TODO: create fallback reasons",
//                 "fallbackUrl": null,
//                 "isFresh": true,
//                 "source": "model",
//                 "model_rank": 3,
//                 "model_score": 0.25,
//                 "page": 0,
//                 "page_rank": 3
//             },
//             "publication": null,
//             "post": null,
//             "comment": {
//                 "name": "Lukas",
//                 "handle": "lukasxp",
//                 "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/0bdedf29-6b37-48a9-ad99-683b67cd28f8_400x400.png",
//                 "id": 203269782,
//                 "body": "I think the internet is going to become completely unusable until someone invents a cryptographic stamp that validates content was created by a human somehow. My entire X timeline is literally articles written by AI telling you how to use AI to make more AIs, and all the comments are also AI, and I’m worried if I keep seeing this I’ll become AI too",
//                 "body_json": {
//                     "type": "doc",
//                     "attrs": {
//                         "schemaVersion": "v1"
//                     },
//                     "content": [
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "I think the internet is going to become completely unusable until someone invents a cryptographic stamp that validates content was created by a human somehow. My entire X timeline is literally articles written by AI telling you how to use AI to make more AIs, and all the comments are also AI, and I’m worried if I keep seeing this I’ll become AI too"
//                                 }
//                             ]
//                         }
//                     ]
//                 },
//                 "publication_id": null,
//                 "post_id": null,
//                 "user_id": 40229043,
//                 "type": "feed",
//                 "date": "2026-01-22T04:02:52.199Z",
//                 "edited_at": null,
//                 "ancestor_path": "",
//                 "reply_minimum_role": "everyone",
//                 "media_clip_id": null,
//                 "reaction_count": 1093,
//                 "reactions": {
//                     "❤": 1093
//                 },
//                 "restacks": 78,
//                 "restacked": false,
//                 "children_count": 120,
//                 "attachments": [],
//                 "user_bestseller_tier": null,
//                 "userStatus": {
//                     "bestsellerTier": null,
//                     "subscriberTier": null,
//                     "leaderboard": null,
//                     "vip": false,
//                     "badge": null,
//                     "paidPublicationIds": [],
//                     "subscriber": null
//                 },
//                 "user_primary_publication": {
//                     "id": 389062,
//                     "subdomain": "lukasxp",
//                     "custom_domain_optional": false,
//                     "name": "The Lukas Experience",
//                     "author_id": 40229043,
//                     "user_id": 40229043,
//                     "handles_enabled": false,
//                     "explicit": false,
//                     "is_personal_mode": false,
//                     "payments_state": "enabled",
//                     "pledges_enabled": false
//                 }
//             },
//             "parentComments": [],
//             "canReply": true,
//             "isMuted": false,
//             "trackingParameters": {
//                 "item_primary_entity_key": "c-203269782",
//                 "item_entity_key": "c-203269782",
//                 "item_type": "comment",
//                 "item_comment_id": 203269782,
//                 "item_content_user_id": 40229043,
//                 "item_content_timestamp": "2026-01-22T04:02:52.199Z",
//                 "item_context_type": "comment_restack",
//                 "item_context_type_bucket": "notes",
//                 "item_context_timestamp": "2026-01-22T04:02:52.199Z",
//                 "item_context_user_id": 86,
//                 "item_context_user_ids": [
//                     86
//                 ],
//                 "item_can_reply": true,
//                 "item_is_fresh": true,
//                 "item_last_impression_at": null,
//                 "item_source": "model",
//                 "item_model_rank": 3,
//                 "item_model_score": 0.25,
//                 "source_comment_id": 209255871,
//                 "item_page": 0,
//                 "item_page_rank": 3,
//                 "tab_id": "for-you",
//                 "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//                 "impression_id": "f31f435b-8cd1-45d2-96fc-3060fb4b239c",
//                 "followed_user_count": 224,
//                 "subscribed_publication_count": 205,
//                 "is_following": false,
//                 "is_explicitly_subscribed": false,
//                 "note_velocity_factor": 1.078759961279,
//                 "note_delay_seconds": 141,
//                 "note_notes_per_hour": 2819.426967,
//                 "item_current_reaction_count": 1093,
//                 "item_current_restack_count": 78,
//                 "item_current_reply_count": 120
//             },
//             "canShowFollowUpsell": false
//         },
//         {
//             "entity_key": "c-210162608",
//             "type": "comment",
//             "context": {
//                 "type": "note",
//                 "typeBucket": "notes",
//                 "timestamp": "2026-02-05T05:05:36.748Z",
//                 "users": [
//                     {
//                         "id": 63039745,
//                         "name": "Daniel Muñoz",
//                         "handle": "bigifftrue",
//                         "previous_name": "Big Iff True",
//                         "photo_url": "https://substackcdn.com/image/fetch/$s_!6boI!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F7cf94bc9-5cb0-40a9-9afe-6378db2c402c_1336x1336.jpeg",
//                         "bio": "I write about philosophy, politics, and economics at Big iff True.",
//                         "profile_set_up_at": "2023-04-08T03:34:30.056Z",
//                         "reader_installed_at": "2025-05-05T16:22:27.133Z",
//                         "bestseller_tier": null,
//                         "status": {
//                             "bestsellerTier": null,
//                             "subscriberTier": 5,
//                             "leaderboard": null,
//                             "vip": false,
//                             "badge": {
//                                 "type": "subscriber",
//                                 "tier": 5,
//                                 "accent_colors": null
//                             },
//                             "paidPublicationIds": [
//                                 375183,
//                                 2880588,
//                                 159185,
//                                 5247799,
//                                 938945,
//                                 876842
//                             ],
//                             "subscriber": null
//                         },
//                         "primary_publication": {
//                             "id": 2883774,
//                             "subdomain": "bigifftrue",
//                             "custom_domain_optional": false,
//                             "name": "Big iff True",
//                             "logo_url": "https://substack-post-media.s3.amazonaws.com/public/images/67b5fd25-9d90-44b7-9b44-8fdbb8bf88a8_1024x1024.png",
//                             "author_id": 63039745,
//                             "user_id": 63039745,
//                             "handles_enabled": false,
//                             "explicit": false,
//                             "is_personal_mode": false,
//                             "payments_state": "disabled",
//                             "pledges_enabled": false
//                         }
//                     }
//                 ],
//                 "fallbackReason": "//TODO: create fallback reasons",
//                 "fallbackUrl": null,
//                 "isFresh": true,
//                 "source": "model",
//                 "model_rank": 4,
//                 "model_score": 0.2,
//                 "page": 0,
//                 "page_rank": 4
//             },
//             "publication": null,
//             "post": null,
//             "comment": {
//                 "name": "Daniel Muñoz",
//                 "handle": "bigifftrue",
//                 "photo_url": "https://substackcdn.com/image/fetch/$s_!6boI!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F7cf94bc9-5cb0-40a9-9afe-6378db2c402c_1336x1336.jpeg",
//                 "id": 210162608,
//                 "body": "An almost certainly fake journal just asked me to review a paper with the following extraordinary abstract, which I still can’t believe I’m reading.\n\n\n\nHuman appetite is so powerful that it leaves no one indifferent; for this very reason, it is possible to reflect again on topics that have been considered settled, such as duty within human relationships, or not. The square of a negative number. The constitution of the real thing as an absolute and particular whole. The search to update the structure of our social reality as a vice versa within our relationships. One cannot exist without the other, which we call p ^ ¬ p. Furthermore, if we speak of two elements or peoples, it is because there are 0, 1, and 01 (¬x, x, x ^¬x). I cannot exist without the others.",
//                 "body_json": {
//                     "type": "doc",
//                     "attrs": {
//                         "schemaVersion": "v1"
//                     },
//                     "content": [
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "An almost certainly fake journal just asked me to review a paper with the following extraordinary abstract, which I still can’t believe I’m reading."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "blockquote",
//                             "content": [
//                                 {
//                                     "type": "paragraph",
//                                     "content": [
//                                         {
//                                             "type": "text",
//                                             "text": "Human appetite is so powerful that it leaves no one indifferent; for this very reason, it is possible to reflect again on topics that have been considered settled, such as duty within human relationships, or not. The square of a negative number. The constitution of the real thing as an absolute and particular whole. The search to update the structure of our social reality as a vice versa within our relationships. One cannot exist without the other, which we call p ^ ¬ p. Furthermore, if we speak of two elements or peoples, it is because there are 0, 1, and 01 (¬x, x, x ^¬x). I cannot exist without the others."
//                                         }
//                                     ]
//                                 }
//                             ]
//                         }
//                     ]
//                 },
//                 "publication_id": null,
//                 "post_id": null,
//                 "user_id": 63039745,
//                 "type": "feed",
//                 "date": "2026-02-05T05:05:36.748Z",
//                 "edited_at": null,
//                 "ancestor_path": "",
//                 "reply_minimum_role": "everyone",
//                 "media_clip_id": null,
//                 "reaction_count": 92,
//                 "reactions": {
//                     "❤": 92
//                 },
//                 "restacks": 9,
//                 "restacked": false,
//                 "children_count": 24,
//                 "attachments": [
//                     {
//                         "id": "18faabd5-1c05-4677-8a0a-fe164c8e53b0",
//                         "type": "image",
//                         "imageUrl": "https://substack-post-media.s3.amazonaws.com/public/images/864e956f-2ace-4527-945d-12054d5884d8_600x400.png",
//                         "imageWidth": 600,
//                         "imageHeight": 400,
//                         "explicit": false
//                     }
//                 ],
//                 "user_bestseller_tier": null,
//                 "userStatus": {
//                     "bestsellerTier": null,
//                     "subscriberTier": 5,
//                     "leaderboard": null,
//                     "vip": false,
//                     "badge": {
//                         "type": "subscriber",
//                         "tier": 5,
//                         "accent_colors": null
//                     },
//                     "paidPublicationIds": [
//                         375183,
//                         2880588,
//                         159185,
//                         5247799,
//                         938945,
//                         876842
//                     ],
//                     "subscriber": null
//                 },
//                 "user_primary_publication": {
//                     "id": 2883774,
//                     "subdomain": "bigifftrue",
//                     "custom_domain_optional": false,
//                     "name": "Big iff True",
//                     "logo_url": "https://substack-post-media.s3.amazonaws.com/public/images/67b5fd25-9d90-44b7-9b44-8fdbb8bf88a8_1024x1024.png",
//                     "author_id": 63039745,
//                     "user_id": 63039745,
//                     "handles_enabled": false,
//                     "explicit": false,
//                     "is_personal_mode": false,
//                     "payments_state": "disabled",
//                     "pledges_enabled": false
//                 }
//             },
//             "parentComments": [],
//             "canReply": true,
//             "isMuted": false,
//             "trackingParameters": {
//                 "item_primary_entity_key": "c-210162608",
//                 "item_entity_key": "c-210162608",
//                 "item_type": "comment",
//                 "item_comment_id": 210162608,
//                 "item_content_user_id": 63039745,
//                 "item_content_timestamp": "2026-02-05T05:05:36.748Z",
//                 "item_context_type": "note",
//                 "item_context_type_bucket": "notes",
//                 "item_context_timestamp": "2026-02-05T05:05:36.748Z",
//                 "item_context_user_id": 63039745,
//                 "item_context_user_ids": [
//                     63039745
//                 ],
//                 "item_can_reply": true,
//                 "item_is_fresh": true,
//                 "item_last_impression_at": null,
//                 "item_source": "model",
//                 "item_model_rank": 4,
//                 "item_model_score": 0.2,
//                 "source_comment_id": 209255871,
//                 "item_page": 0,
//                 "item_page_rank": 4,
//                 "tab_id": "for-you",
//                 "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//                 "impression_id": "8701d7d8-d8c5-48ba-8eaa-2ae5124ebd2e",
//                 "followed_user_count": 224,
//                 "subscribed_publication_count": 205,
//                 "is_following": false,
//                 "is_explicitly_subscribed": false,
//                 "note_velocity_factor": 1.078759961279,
//                 "note_delay_seconds": 141,
//                 "note_notes_per_hour": 2819.426967,
//                 "item_current_reaction_count": 92,
//                 "item_current_restack_count": 9,
//                 "item_current_reply_count": 24
//             },
//             "canShowFollowUpsell": false
//         },
//         {
//             "entity_key": "c-210294224",
//             "type": "comment",
//             "context": {
//                 "type": "note",
//                 "typeBucket": "notes",
//                 "timestamp": "2026-02-05T14:20:39.146Z",
//                 "users": [
//                     {
//                         "id": 6831253,
//                         "name": "Luiza Jarovsky, PhD",
//                         "handle": "luizajarovsky",
//                         "previous_name": "Luiza Jarovsky",
//                         "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/2f2904b3-0f07-4359-a28a-cca400d254e8_1604x1604.jpeg",
//                         "bio": "Co-founder of the AI, Tech & Privacy Academy (1,400+ participants), Author of Luiza's Newsletter (90,000+ subscribers), Mother of 3.",
//                         "profile_set_up_at": "2022-03-20T23:50:55.278Z",
//                         "reader_installed_at": "2024-08-21T21:38:36.009Z",
//                         "bestseller_tier": 100,
//                         "status": {
//                             "bestsellerTier": 100,
//                             "subscriberTier": null,
//                             "leaderboard": {
//                                 "ranking": "paid",
//                                 "rank": 54,
//                                 "publicationName": "Luiza's Newsletter",
//                                 "label": "Technology",
//                                 "categoryId": "4",
//                                 "publicationId": 808767
//                             },
//                             "vip": false,
//                             "badge": {
//                                 "type": "bestseller",
//                                 "tier": 100
//                             },
//                             "paidPublicationIds": [],
//                             "subscriber": null
//                         },
//                         "primary_publication": {
//                             "id": 808767,
//                             "subdomain": "luizasnewsletter",
//                             "custom_domain": "www.luizasnewsletter.com",
//                             "custom_domain_optional": false,
//                             "name": "Luiza's Newsletter",
//                             "logo_url": "https://substack-post-media.s3.amazonaws.com/public/images/b6a56d5a-1724-4dac-b72a-afd9e41051e9_600x600.png",
//                             "author_id": 6831253,
//                             "user_id": 6831253,
//                             "handles_enabled": false,
//                             "explicit": false,
//                             "is_personal_mode": false,
//                             "payments_state": "enabled",
//                             "pledges_enabled": false
//                         }
//                     }
//                 ],
//                 "fallbackReason": "//TODO: create fallback reasons",
//                 "fallbackUrl": null,
//                 "isFresh": true,
//                 "source": "model",
//                 "model_rank": 5,
//                 "model_score": 0.16666666666666666,
//                 "page": 0,
//                 "page_rank": 5
//             },
//             "publication": null,
//             "post": null,
//             "comment": {
//                 "name": "Luiza Jarovsky, PhD",
//                 "handle": "luizajarovsky",
//                 "photo_url": "https://substack-post-media.s3.amazonaws.com/public/images/2f2904b3-0f07-4359-a28a-cca400d254e8_1604x1604.jpeg",
//                 "id": 210294224,
//                 "body": "Anthropic’s ads mocking ChatGPT ads (and sycophancy) are funny, but also hypocritical.\n\nThis is the same company that published a highly anthropomorphic \"constitution\" for its AI model, quietly diluting ideas of legal liability and fundamental rights (here: https://www.luizasnewsletter.com/p/claudes-strange-constitution).\n\nWe are paying attention.",
//                 "body_json": {
//                     "type": "doc",
//                     "attrs": {
//                         "schemaVersion": "v1"
//                     },
//                     "content": [
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "Anthropic’s ads mocking ChatGPT ads (and sycophancy) are funny, but also hypocritical."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "This is the same company that published a highly anthropomorphic \"constitution\" for its AI model, quietly diluting ideas of legal liability and fundamental rights (here: "
//                                 },
//                                 {
//                                     "type": "text",
//                                     "marks": [
//                                         {
//                                             "type": "link",
//                                             "attrs": {
//                                                 "href": "https://www.luizasnewsletter.com/p/claudes-strange-constitution",
//                                                 "target": "_blank",
//                                                 "rel": "nofollow ugc noopener",
//                                                 "class": "note-link"
//                                             }
//                                         }
//                                     ],
//                                     "text": "https://www.luizasnewsletter.com/p/claudes-strange-constitution"
//                                 },
//                                 {
//                                     "type": "text",
//                                     "text": ")."
//                                 }
//                             ]
//                         },
//                         {
//                             "type": "paragraph",
//                             "content": [
//                                 {
//                                     "type": "text",
//                                     "text": "We are paying attention."
//                                 }
//                             ]
//                         }
//                     ]
//                 },
//                 "publication_id": null,
//                 "post_id": null,
//                 "user_id": 6831253,
//                 "type": "feed",
//                 "date": "2026-02-05T14:20:39.146Z",
//                 "edited_at": null,
//                 "ancestor_path": "",
//                 "reply_minimum_role": "everyone",
//                 "media_clip_id": null,
//                 "reaction_count": 43,
//                 "reactions": {
//                     "❤": 43
//                 },
//                 "restacks": 5,
//                 "restacked": false,
//                 "children_count": 10,
//                 "attachments": [
//                     {
//                         "id": "b8a87aed-16c2-4759-8a7a-68c173e1cfe4",
//                         "user_id": 6831253,
//                         "comment_id": 210294224,
//                         "type": "video",
//                         "media_upload_id": "c7571101-9f55-4329-9e6b-96bc750de1ac",
//                         "mediaUpload": {
//                             "id": "c7571101-9f55-4329-9e6b-96bc750de1ac",
//                             "name": "lBdPDZgke9UoWlyA.mp4",
//                             "created_at": "2026-02-05T14:18:16.893Z",
//                             "uploaded_at": "2026-02-05T14:18:19.059Z",
//                             "publication_id": null,
//                             "state": "transcoded",
//                             "post_id": null,
//                             "user_id": 6831253,
//                             "duration": 60.166668,
//                             "height": 720,
//                             "width": 1280,
//                             "thumbnail_id": 1,
//                             "preview_start": null,
//                             "preview_duration": null,
//                             "media_type": "video",
//                             "primary_file_size": 3696495,
//                             "is_mux": true,
//                             "mux_asset_id": "aYDbkPmyJ3vdHOiknhIBIOCte3i01b01pP8qJ1b2U8S028",
//                             "mux_playback_id": "tNz5JJ4tf015rV74yVsj2mcGtPO3rr02tmdM4jKr9nTkw",
//                             "mux_preview_asset_id": null,
//                             "mux_preview_playback_id": null,
//                             "mux_rendition_quality": "high",
//                             "mux_preview_rendition_quality": null,
//                             "explicit": false,
//                             "copyright_infringement": null,
//                             "src_media_upload_id": null,
//                             "live_stream_id": null
//                         }
//                     }
//                 ],
//                 "user_bestseller_tier": 100,
//                 "userStatus": {
//                     "bestsellerTier": 100,
//                     "subscriberTier": null,
//                     "leaderboard": {
//                         "ranking": "paid",
//                         "rank": 54,
//                         "publicationName": "Luiza's Newsletter",
//                         "label": "Technology",
//                         "categoryId": "4",
//                         "publicationId": 808767
//                     },
//                     "vip": false,
//                     "badge": {
//                         "type": "bestseller",
//                         "tier": 100
//                     },
//                     "paidPublicationIds": [],
//                     "subscriber": null
//                 },
//                 "user_primary_publication": {
//                     "id": 808767,
//                     "subdomain": "luizasnewsletter",
//                     "custom_domain": "www.luizasnewsletter.com",
//                     "custom_domain_optional": false,
//                     "name": "Luiza's Newsletter",
//                     "logo_url": "https://substack-post-media.s3.amazonaws.com/public/images/b6a56d5a-1724-4dac-b72a-afd9e41051e9_600x600.png",
//                     "author_id": 6831253,
//                     "user_id": 6831253,
//                     "handles_enabled": false,
//                     "explicit": false,
//                     "is_personal_mode": false,
//                     "payments_state": "enabled",
//                     "pledges_enabled": false
//                 }
//             },
//             "parentComments": [],
//             "canReply": true,
//             "isMuted": false,
//             "trackingParameters": {
//                 "item_primary_entity_key": "c-210294224",
//                 "item_entity_key": "c-210294224",
//                 "item_type": "comment",
//                 "item_comment_id": 210294224,
//                 "item_content_user_id": 6831253,
//                 "item_content_timestamp": "2026-02-05T14:20:39.146Z",
//                 "item_context_type": "note",
//                 "item_context_type_bucket": "notes",
//                 "item_context_timestamp": "2026-02-05T14:20:39.146Z",
//                 "item_context_user_id": 6831253,
//                 "item_context_user_ids": [
//                     6831253
//                 ],
//                 "item_can_reply": true,
//                 "item_is_fresh": true,
//                 "item_last_impression_at": null,
//                 "item_source": "model",
//                 "item_model_rank": 5,
//                 "item_model_score": 0.16666666666666666,
//                 "source_comment_id": 209255871,
//                 "item_page": 0,
//                 "item_page_rank": 5,
//                 "tab_id": "for-you",
//                 "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//                 "impression_id": "a3541056-ebe8-4259-89e2-278b8776c72e",
//                 "followed_user_count": 224,
//                 "subscribed_publication_count": 205,
//                 "is_following": false,
//                 "is_explicitly_subscribed": false,
//                 "note_velocity_factor": 1.078759961279,
//                 "note_delay_seconds": 141,
//                 "note_notes_per_hour": 2819.426967,
//                 "item_current_reaction_count": 43,
//                 "item_current_restack_count": 5,
//                 "item_current_reply_count": 10
//             },
//             "canShowFollowUpsell": false
//         }
//     ],
//     "originalCursorTimestamp": "2026-02-06T09:00:58.766Z",
//     "nextCursor": "eyJzZXNzaW9uX2lkIjoiZTIzMjQ4YzUtZjgwMS00MDExLTgzMGUtNWU2NTdkNmFjMDA1IiwiYmVmb3JlX3RpbWVzdGFtcCI6IjIwMjYtMDItMDZUMDk6MDA6NTguNzY2WiIsImFmdGVyX3RpbWVzdGFtcCI6IjIwMjYtMDItMDZUMDk6MDA6NTguNzY2WiIsInBhZ2VfbnVtYmVyIjoxLCJsYXN0X2J1Y2tldF9pdGVtcyI6eyJub3RlcyI6eyJlbnRpdHlfa2V5IjoiYy0yMTAyOTQyMjQiLCJ1c2VyX3Njb3JlIjowLjE2NjY2NjY2NjY2NjY2NjY2fX0sInByZXBlbmRlZF9ub3RlX2lkcyI6W119",
//     "trackingParameters": {
//         "feed_session_id": "e23248c5-f801-4011-830e-5e657d6ac005",
//         "top_note_impression_id": "70a88bfc-3700-479b-b311-47e131acab63",
//         "tab_id": "for-you",
//         "top_note_entity_key": "c-198924049",
//         "top_note_model_score": 1,
//         "top_note_source": "model",
//         "top_note_context_user_ids": [
//             4104627
//         ],
//         "top_note_context_type": "note",
//         "is_following": false,
//         "followed_user_count": 224,
//         "subscribed_publication_count": 205
//     },
//     "tabs": [
//         {
//             "id": "for-you",
//             "name": "For you",
//             "type": "base",
//             "layout": "post_queue_head"
//         },
//         {
//             "id": "subscribed",
//             "name": "Following",
//             "type": "secondary",
//             "layout": "post_queue_head"
//         },
//         {
//             "id": "bestseller",
//             "name": "New Bestsellers",
//             "type": "secondary",
//             "layout": "no_head"
//         },
//         {
//             "id": "96",
//             "name": "Culture",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "culture"
//         },
//         {
//             "id": "4",
//             "name": "Technology",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "technology"
//         },
//         {
//             "id": "62",
//             "name": "Business",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "business"
//         },
//         {
//             "id": "76739",
//             "name": "U.S. Politics",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "us-politics"
//         },
//         {
//             "id": "153",
//             "name": "Finance",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "finance"
//         },
//         {
//             "id": "13645",
//             "name": "Food & Drink",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "food"
//         },
//         {
//             "id": "94",
//             "name": "Sports",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "sports"
//         },
//         {
//             "id": "15417",
//             "name": "Art & Illustration",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "art"
//         },
//         {
//             "id": "76740",
//             "name": "World Politics",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "world-politics"
//         },
//         {
//             "id": "76741",
//             "name": "Health Politics",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "health-politics"
//         },
//         {
//             "id": "49715",
//             "name": "Fashion & Beauty",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "fashionandbeauty"
//         },
//         {
//             "id": "103",
//             "name": "News",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "news"
//         },
//         {
//             "id": "11",
//             "name": "Music",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "music"
//         },
//         {
//             "id": "223",
//             "name": "Faith & Spirituality",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "faith"
//         },
//         {
//             "id": "15414",
//             "name": "Climate & Environment",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "climate"
//         },
//         {
//             "id": "134",
//             "name": "Science",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "science"
//         },
//         {
//             "id": "339",
//             "name": "Literature",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "literature"
//         },
//         {
//             "id": "284",
//             "name": "Fiction",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "fiction"
//         },
//         {
//             "id": "61",
//             "name": "Design",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "design"
//         },
//         {
//             "id": "355",
//             "name": "Health & Wellness",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "health"
//         },
//         {
//             "id": "109",
//             "name": "Travel",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "travel"
//         },
//         {
//             "id": "1796",
//             "name": "Parenting",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "parenting"
//         },
//         {
//             "id": "114",
//             "name": "Philosophy",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "philosophy"
//         },
//         {
//             "id": "387",
//             "name": "Comics",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "comics"
//         },
//         {
//             "id": "51282",
//             "name": "International",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "international"
//         },
//         {
//             "id": "118",
//             "name": "Crypto",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "crypto"
//         },
//         {
//             "id": "18",
//             "name": "History",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "history"
//         },
//         {
//             "id": "49692",
//             "name": "Humor",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "humor"
//         },
//         {
//             "id": "34",
//             "name": "Education",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "education"
//         },
//         {
//             "id": "76782",
//             "name": "Film & TV",
//             "type": "category",
//             "layout": "module_head",
//             "slug": "film-and-tv"
//         }
//     ]
// }
