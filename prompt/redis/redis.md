1. **`hmset`** – Set multiple hash fields to multiple values.

    - Set fields `field1`, `field2`, and `field3` with values `value1`, `value2`, and `value3` for the hash key `my_hash`.

2. **`hget`** – Get the value of a hash field.

    - Get the value of the field `field1` in the hash `my_hash`.

3. **`hgetall`** – Get all the fields and values in a hash.

    - Get all the fields and values in the hash `my_hash`.

4. **`scan`** – Scan Redis keys matching a pattern.

    - Scan for Redis keys that match the pattern `user:*`.

5. **`set`** – Set string value with optional NX and PX options.

    - Set the key `my_key` to the value `my_value` with an expiry of 5000 milliseconds.

6. **`get`** – Get string value.

    - Get the value of the key `my_key`.

7. **`del`** – Delete a key.

    - Delete the key `my_key`.

8. **`zadd`** – Add one or more members to a sorted set.

    - Add members with scores to the sorted set `my_sorted_set`.

9. **`zrange`** – Return a range of members from a sorted set by index.

    - Get members from index 0 to 5 from the sorted set `my_sorted_set` with their scores.

10. **`zrangebyscore`** – Return members from a sorted set with scores between min and max.

    - Get members from the sorted set `my_sorted_set` with scores between 10 and 50.

11. **`zrem`** – Remove one or more members from a sorted set.

    - Remove members `member1` and `member2` from the sorted set `my_sorted_set`.

12. **`sadd`** – Add one or more members to a set.

    - Add members `member1`, `member2`, and `member3` to the set `my_set`.

13. **`smembers`** – Get all members in a set.

    - Get all members in the set `my_set`.
