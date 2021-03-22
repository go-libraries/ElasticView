let a = {
  'number_of_shards':1,
    'index.translog.sync_interval':1,
    'index.shard.check_on_startup':1,
    'index.routing_partition_size':1,
    'index.codec':1,
  'index.shard.check_on_startup': 'false',
  'index.routing_partition_size': 1,
  'index.codec': 'best_compression',
  'index.auto_expand_replicas': 'false',
  'index.max_result_window': 10000,
  'index.max_rescore_window': 10000,
  'index.blocks.read_only': 'false',
  'index.blocks.read': 'false',
  'index.blocks.write': 'false',
  'index.blocks.metadata': 'false',
  'index.blocks.read_only_allow_delete': 'false',
  'index.max_refresh_listeners': 10,
  'number_of_shards': "1",
  'number_of_replicas': 0,
  'refresh_interval': '1s',
  'index.translog.sync_interval': '5s',
  'index.translog.durability': 'async',
  'index.translog.flush_threshold_size': '5g',
  'index.merge.scheduler.max_thread_count': 20,
  'index.merge.policy.max_merged_segment': '5gb',
  'index.merge.policy.segments_per_tier': '10'
}
for(let tmp in a){
  console.log({ caption: tmp, meta: tmp, value: "`" +`"`+tmp.toString()+`"`+ "`"},",")
}
