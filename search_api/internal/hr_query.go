package internal

func CreateQuery(q *Query) map[string]interface{} {
  query := map[string]interface{}{}
  if q.Id != "" {
    query = map[string]interface{}{
      "query": map[string]interface{}{
        "bool": map[string]interface{}{
          "must": []map[string]interface{}{
            {
              "match": map[string]interface{}{
                "referencenumber": q.Id,
              },
            },
          },
        },
      },
    }
  } else if q.Keyword != "" && q.State != "" {
    query = map[string]interface{}{
      "query": map[string]interface{}{
        "bool": map[string]interface{}{
          "must": []map[string]interface{}{
            {
              "bool": map[string]interface{}{
                "should": []map[string]interface{}{
                  {
                    "match": map[string]interface{}{
                      "title": map[string]interface{}{
                        "query": q.Keyword,
                        "boost": 3,
                      },
                    },
                  },
                  {
                    "match": map[string]interface{}{
                      "description": map[string]interface{}{
                        "query": q.Keyword,
                        "boost": 2,
                      },
                    },
                  },
                  {
                    "match": map[string]interface{}{
                      "category": map[string]interface{}{
                        "query": q.Keyword,
                        "boost": 1,
                      },
                    },
                  },
                },
                "minimum_should_match": 1,
              },
            },
            {
              "bool": map[string]interface{}{
                "must": []map[string]interface{}{
                  {
                    "match": map[string]interface{}{
                      "state": q.State,
                    },
                  },
                },
              },
            },
          },
        },
      },
    }
  } else if q.Keyword != "" && q.State == "" {
    query = map[string]interface{}{
      "query": map[string]interface{}{
        "bool": map[string]interface{}{
          "should": []map[string]interface{}{
            {
              "match": map[string]interface{}{
                "title": map[string]interface{}{
                  "query": q.Keyword,
                  "boost": 3,
                },
              },
            },
            {
              "match": map[string]interface{}{
                "description": map[string]interface{}{
                  "query": q.Keyword,
                  "boost": 2,
                },
              },
            },
            {
              "match": map[string]interface{}{
                "category": map[string]interface{}{
                  "query": q.Keyword,
                  "boost": 1,
                },
              },
            },
          },
          "minimum_should_match": 1,
        },
      },
    }
  } else if q.Keyword == "" && q.State != "" {
    query = map[string]interface{}{
      "query": map[string]interface{}{
        "bool": map[string]interface{}{
          "must": []map[string]interface{}{
            {
              "match": map[string]interface{}{
                "state": q.State,
              },
            },
          },
        },
      },
    }
  }
  return query
}