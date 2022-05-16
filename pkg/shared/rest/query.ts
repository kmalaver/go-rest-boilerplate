const params = {
  filter: {
    name: { _in: ['a', 'b', 'c'] },
    age: { _gt: 20 },
    _or: {
      role: { _in: ['admin', 'user'] },
      has_permission: true
    }
  },
  order_by: {
    name: 'asc',
  },
  limit: 10,
  offset: 0,
}

// encode above params to query string