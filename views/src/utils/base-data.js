export const esPathKeyWords = [
  { 'data': '_cat/indices', 'value': '查看所有索引 _cat/indices' },
  { 'data': '_cat/master', 'value': '查看集群master节点 _cat/master' },
  { 'data': '_cat/health', 'value': '查看集群健康状态  _cat/health' },
  { 'data': '_cat/nodes', 'value': '查看集群节点和磁盘剩余  _cat/nodes' },
  { 'data': '_cat/allocation', 'value': '查看分配  _cat/allocation' },
  { 'data': '_cat/pending_tasks', 'value': '查看被挂起任务  _cat/pending_tasks' },
  { 'data': '_cat/plugins', 'value': '查看每个节点正在运行的插件  _cat/plugins' },
  { 'data': '_cat/nodeattrs', 'value': '查看每个节点的自定义属性  _cat/nodeattrs' },
  { 'data': '_cat/recovery', 'value': '查看索引分片的恢复视图 _cat/recovery' },
  { 'data': '_cat/fielddata', 'value': '查看每个数据节点上fielddata当前占用的堆内存 _cat/fielddata' },
  { 'data': '_cat/repositories', 'value': '查看注册的快照仓库 _cat/repositories' },
  { 'data': '_cat/snapshots', 'value': '查看快照仓库下的快照 _cat/snapshots' },
  { 'data': '_cat/thread_pool/bulk', 'value': '查看每个节点线程池的统计信息 _cat/thread_pool/bulk' },
  { 'data': '_cat/aliases', 'value': '查看别名 _cat/aliases' },
  { 'data': '_cat/templates', 'value': '查看索引模板 _cat/templates' },
  { 'data': '_cat/count', 'value': '查看单个或某类或整个集群文档数 _cat/count' },
  { 'data': '_cat/shards', 'value': '查看每个索引的分片 _cat/shards' },
  { 'data': '_cat/segments', 'value': '查看每个索引的segment _cat/segments' },
  { 'data': '_bulk', 'value': '_bulk' },
  { 'data': '{indices}/_bulk', 'value': '{indices}/_bulk' },
  { 'data': '{indices}/{type}/_bulk', 'value': '{indices}/{type}/_bulk' },
  { 'data': '_cat/aliases/{name}', 'value': '_cat/aliases/{name}' },
  { 'data': '_cat/allocation/{nodes}', 'value': '_cat/allocation/{nodes}' },
  { 'data': '_cat/count/{indices}', 'value': '_cat/count/{indices}' },
  { 'data': '_cat/fielddata/{fields}', 'value': '_cat/fielddata/{fields}' },
  { 'data': '_cat', 'value': '_cat' },
  { 'data': '_cat/indices/{indices}', 'value': '_cat/indices/{indices}' },
  { 'data': '_cat/recovery/{indices}', 'value': '_cat/recovery/{indices}' },
  { 'data': '_cat/segments/{indices}', 'value': '_cat/segments/{indices}' },
  { 'data': '_cat/shards/{indices}', 'value': '_cat/shards/{indices}' },
  { 'data': '_cat/snapshots/{repository}', 'value': '_cat/snapshots/{repository}' },
  { 'data': '_cat/tasks', 'value': '_cat/tasks' },
  { 'data': '_cat/templates/{name}', 'value': '_cat/templates/{name}' },
  { 'data': '_cat/thread_pool', 'value': '_cat/thread_pool' },
  { 'data': '_cat/thread_pool/{thread_pool_patterns}', 'value': '_cat/thread_pool/{thread_pool_patterns}' },
  { 'data': '_search/scroll/{scroll_id}', 'value': '_search/scroll/{scroll_id}' },
  { 'data': '_search/scroll', 'value': '_search/scroll' },
  { 'data': '_cluster/allocation/explain', 'value': '_cluster/allocation/explain' },
  { 'data': '_cluster/settings', 'value': '_cluster/settings' },
  { 'data': '_cluster/health', 'value': '_cluster/health' },
  { 'data': '_cluster/health/{indices}', 'value': '_cluster/health/{indices}' },
  { 'data': '_cluster/pending_tasks', 'value': '_cluster/pending_tasks' },
  { 'data': '_remote/info', 'value': '_remote/info' },
  { 'data': '_cluster/reroute', 'value': '_cluster/reroute' },
  { 'data': '_cluster/state', 'value': '_cluster/state' },
  { 'data': '_cluster/state/{metrics}', 'value': '_cluster/state/{metrics}' },
  { 'data': '_cluster/state/{metrics}/{indices}', 'value': '_cluster/state/{metrics}/{indices}' },
  { 'data': '_cluster/stats', 'value': '_cluster/stats' },
  { 'data': '_cluster/stats/nodes/{nodes}', 'value': '_cluster/stats/nodes/{nodes}' },
  { 'data': '_count', 'value': '_count' },
  { 'data': '{indices}/_count', 'value': '{indices}/_count' },
  { 'data': '{indices}/{type}/_count', 'value': '{indices}/{type}/_count' },
  { 'data': '{indices}/{type}/{id}/_create', 'value': '{indices}/{type}/{id}/_create' },
  { 'data': '{indices}/_delete_by_query', 'value': '{indices}/_delete_by_query' },
  { 'data': '{indices}/{type}/_delete_by_query', 'value': '{indices}/{type}/_delete_by_query' },
  { 'data': '_scripts/{id}', 'value': '_scripts/{id}' },
  { 'data': '{indices}/{type}/{id}', 'value': '{indices}/{type}/{id}' },
  { 'data': '{indices}/_doc/{id}', 'value': '{indices}/_doc/{id}' },
  { 'data': '{indices}/{type}/{id}/_source', 'value': '{indices}/{type}/{id}/_source' },
  { 'data': '{indices}/{type}/{id}/_explain', 'value': '{indices}/{type}/{id}/_explain' },
  {
    'data': '_field_caps',
    'value': '_field_caps'
  },
  { 'data': '{indices}/_field_caps', 'value': '{indices}/_field_caps' },
  {
    'data': '{indices}/{type}',
    'value': '{indices}/{type}'
  },
  { 'data': '{indices}/_doc', 'value': '{indices}/_doc' },
  {
    'data': '_analyze',
    'value': '_analyze'
  },
  { 'data': '{indices}/_analyze', 'value': '{indices}/_analyze' },
  {
    'data': '_cache/clear',
    'value': '_cache/clear'
  },
  { 'data': '{indices}/_cache/clear', 'value': '{indices}/_cache/clear' },
  {
    'data': '{indices}/_close',
    'value': '{indices}/_close'
  },
  { 'data': '{indices}', 'value': '{indices}' },
  {
    'data': '{indices}/_alias/{name}',
    'value': '{indices}/_alias/{name}'
  },
  { 'data': '{indices}/_aliases/{name}', 'value': '{indices}/_aliases/{name}' },
  {
    'data': '_template/{template}',
    'value': '_template/{template}'
  },
  { 'data': '_alias/{name}', 'value': '_alias/{name}' },
  {
    'data': '{indices}/_mapping/{type}',
    'value': '{indices}/_mapping/{type}'
  },
  { 'data': '_flush/synced', 'value': '_flush/synced' },
  {
    'data': '{indices}/_flush/synced',
    'value': '{indices}/_flush/synced'
  },
  { 'data': '_flush', 'value': '_flush' },
  {
    'data': '{indices}/_flush',
    'value': '{indices}/_flush'
  },
  { 'data': '_forcemerge', 'value': '_forcemerge' },
  {
    'data': '{indices}/_forcemerge',
    'value': '{indices}/_forcemerge'
  },
  { 'data': '_alias', 'value': '_alias' },
  {
    'data': '{indices}/_alias',
    'value': '{indices}/_alias'
  },
  {
    'data': '_mapping/field/{fields}',
    'value': '_mapping/field/{fields}'
  },
  {
    'data': '{indices}/_mapping/field/{fields}',
    'value': '{indices}/_mapping/field/{fields}'
  },
  {
    'data': '_mapping/{type}/field/{fields}',
    'value': '_mapping/{type}/field/{fields}'
  },
  {
    'data': '{indices}/_mapping/{type}/field/{fields}',
    'value': '{indices}/_mapping/{type}/field/{fields}'
  },
  { 'data': '_mapping', 'value': '_mapping' },
  {
    'data': '{indices}/_mapping',
    'value': '{indices}/_mapping'
  },
  { 'data': '_mapping/{type}', 'value': '_mapping/{type}' },
  {
    'data': '_settings',
    'value': '_settings'
  },
  { 'data': '{indices}/_settings', 'value': '{indices}/_settings' },
  {
    'data': '{indices}/_settings/{name}',
    'value': '{indices}/_settings/{name}'
  },
  { 'data': '_settings/{name}', 'value': '_settings/{name}' },
  {
    'data': '_template',
    'value': '_template'
  },
  { 'data': '_upgrade', 'value': '_upgrade' },
  {
    'data': '{indices}/_upgrade',
    'value': '{indices}/_upgrade'
  },
  { 'data': '{indices}/_open', 'value': '{indices}/_open' },
  {
    'data': '{indices}/{type}/_mapping',
    'value': '{indices}/{type}/_mapping'
  },
  { 'data': '{indices}/{type}/_mappings', 'value': '{indices}/{type}/_mappings' },
  {
    'data': '{indices}/_mappings/{type}',
    'value': '{indices}/_mappings/{type}'
  },
  { 'data': '_mappings/{type}', 'value': '_mappings/{type}' },
  {
    'data': '{indices}/_mappings',
    'value': '{indices}/_mappings'
  },
  { 'data': '_recovery', 'value': '_recovery' },
  {
    'data': '{indices}/_recovery',
    'value': '{indices}/_recovery'
  },
  { 'data': '_refresh', 'value': '_refresh' },
  {
    'data': '{indices}/_refresh',
    'value': '{indices}/_refresh'
  },
  { 'data': '{alias}/_rollover', 'value': '{alias}/_rollover' },
  {
    'data': '{alias}/_rollover/{new_index}',
    'value': '{alias}/_rollover/{new_index}'
  },
  { 'data': '_segments', 'value': '_segments' },
  {
    'data': '{indices}/_segments',
    'value': '{indices}/_segments'
  },
  { 'data': '_shard_stores', 'value': '_shard_stores' },
  {
    'data': '{indices}/_shard_stores',
    'value': '{indices}/_shard_stores'
  },
  { 'data': '{indices}/_shrink/{target}', 'value': '{indices}/_shrink/{target}' },
  {
    'data': '{indices}/_split/{target}',
    'value': '{indices}/_split/{target}'
  },
  { 'data': '_stats', 'value': '_stats' },
  {
    'data': '_stats/{metrics}',
    'value': '_stats/{metrics}'
  },
  { 'data': '{indices}/_stats', 'value': '{indices}/_stats' },
  {
    'data': '{indices}/_stats/{metrics}',
    'value': '{indices}/_stats/{metrics}'
  },
  { 'data': '_aliases', 'value': '_aliases' },
  {
    'data': '_validate/query',
    'value': '_validate/query'
  },
  {
    'data': '{indices}/_validate/query',
    'value': '{indices}/_validate/query'
  },
  { 'data': '{indices}/{type}/_validate/query', 'value': '{indices}/{type}/_validate/query' },
  { 'data': '_ingest/pipeline/{id}', 'value': '_ingest/pipeline/{id}' },
  {
    'data': '_ingest/pipeline',
    'value': '_ingest/pipeline'
  },
  { 'data': '_ingest/processor/grok', 'value': '_ingest/processor/grok' },
  {
    'data': '_ingest/pipeline/_simulate',
    'value': '_ingest/pipeline/_simulate'
  },
  { 'data': '_ingest/pipeline/{id}/_simulate', 'value': '_ingest/pipeline/{id}/_simulate' },
  {
    'data': '_mget',
    'value': '_mget'
  },
  { 'data': '{indices}/_mget', 'value': '{indices}/_mget' },
  {
    'data': '{indices}/{type}/_mget',
    'value': '{indices}/{type}/_mget'
  },
  { 'data': '_msearch/template', 'value': '_msearch/template' },
  {
    'data': '{indices}/_msearch/template',
    'value': '{indices}/_msearch/template'
  },
  { 'data': '{indices}/{type}/_msearch/template', 'value': '{indices}/{type}/_msearch/template' },
  {
    'data': '_msearch',
    'value': '_msearch'
  },
  { 'data': '{indices}/_msearch', 'value': '{indices}/_msearch' },
  {
    'data': '{indices}/{type}/_msearch',
    'value': '{indices}/{type}/_msearch'
  },
  { 'data': '_mtermvectors', 'value': '_mtermvectors' },
  {
    'data': '{indices}/_mtermvectors',
    'value': '{indices}/_mtermvectors'
  },
  {
    'data': '{indices}/{type}/_mtermvectors',
    'value': '{indices}/{type}/_mtermvectors'
  },
  { 'data': '_cluster/nodes/hotthreads', 'value': '_cluster/nodes/hotthreads' },

  {
    'data': '_all/_settings',
    'value': '_all/_settings'
  },
  {
    'data': '_cluster/nodes/hot_threads',
    'value': '_cluster/nodes/hot_threads'
  },
  {
    'data': '_cluster/nodes/{nodes}/hotthreads',
    'value': '_cluster/nodes/{nodes}/hotthreads'
  },
  {
    'data': '_cluster/nodes/{nodes}/hot_threads',
    'value': '_cluster/nodes/{nodes}/hot_threads'
  },
  { 'data': '_nodes/hotthreads', 'value': '_nodes/hotthreads' },
  {
    'data': '_nodes/hot_threads',
    'value': '_nodes/hot_threads'
  },
  { 'data': '_nodes/{nodes}/hotthreads', 'value': '_nodes/{nodes}/hotthreads' },
  {
    'data': '_nodes/{nodes}/hot_threads',
    'value': '_nodes/{nodes}/hot_threads'
  },
  { 'data': '_nodes', 'value': '_nodes' },
  {
    'data': '_nodes/{nodes}',
    'value': '_nodes/{nodes}'
  },
  { 'data': '_nodes/{metrics}', 'value': '_nodes/{metrics}' },
  {
    'data': '_nodes/{nodes}/{metrics}',
    'value': '_nodes/{nodes}/{metrics}'
  },
  { 'data': '_nodes/stats', 'value': '_nodes/stats' },
  {
    'data': '_nodes/{nodes}/stats',
    'value': '_nodes/{nodes}/stats'
  },
  { 'data': '_nodes/stats/{metrics}', 'value': '_nodes/stats/{metrics}' },
  {
    'data': '_nodes/{nodes}/stats/{metrics}',
    'value': '_nodes/{nodes}/stats/{metrics}'
  },
  {
    'data': '_nodes/stats/{metrics}/{index_metric}',
    'value': '_nodes/stats/{metrics}/{index_metric}'
  },
  {
    'data': '_nodes/{nodes}/stats/{metrics}/{index_metric}',
    'value': '_nodes/{nodes}/stats/{metrics}/{index_metric}'
  },
  { 'data': '_nodes/usage', 'value': '_nodes/usage' },
  {
    'data': '_nodes/{nodes}/usage',
    'value': '_nodes/{nodes}/usage'
  },
  { 'data': '_nodes/usage/{metrics}', 'value': '_nodes/usage/{metrics}' },
  {
    'data': '_nodes/{nodes}/usage/{metrics}',
    'value': '_nodes/{nodes}/usage/{metrics}'
  },
  { 'data': '_scripts/{lang}/{id}', 'value': '_scripts/{lang}/{id}' },
  {
    'data': '_scripts/{lang}/{id}/_create',
    'value': '_scripts/{lang}/{id}/_create'
  },
  { 'data': '_rank_eval', 'value': '_rank_eval' },
  {
    'data': '{indices}/_rank_eval',
    'value': '{indices}/_rank_eval'
  },
  {
    'data': '_reindex/{task_id}/_rethrottle',
    'value': '_reindex/{task_id}/_rethrottle'
  },
  {
    'data': '_update_by_query/{task_id}/_rethrottle',
    'value': '_update_by_query/{task_id}/_rethrottle'
  },
  {
    'data': '_delete_by_query/{task_id}/_rethrottle',
    'value': '_delete_by_query/{task_id}/_rethrottle'
  },
  { 'data': '_reindex', 'value': '_reindex' },
  {
    'data': '_render/template',
    'value': '_render/template'
  },
  { 'data': '_render/template/{id}', 'value': '_render/template/{id}' },
  {
    'data': '_scripts/painless/_execute',
    'value': '_scripts/painless/_execute'
  },
  { 'data': '_search_shards', 'value': '_search_shards' },
  {
    'data': '{indices}/_search_shards',
    'value': '{indices}/_search_shards'
  },
  { 'data': '_search/template', 'value': '_search/template' },
  {
    'data': '{indices}/_search/template',
    'value': '{indices}/_search/template'
  },
  { 'data': '{indices}/{type}/_search/template', 'value': '{indices}/{type}/_search/template' },
  {
    'data': '_search',
    'value': '_search'
  },
  { 'data': '{indices}/_search', 'value': '{indices}/_search' },
  {
    'data': '{indices}/{type}/_search',
    'value': '{indices}/{type}/_search'
  },
  { 'data': '_snapshot/{repository}', 'value': '_snapshot/{repository}' },
  {
    'data': '_snapshot/{repository}/{snapshot}',
    'value': '_snapshot/{repository}/{snapshot}'
  },
  { 'data': '_snapshot', 'value': '_snapshot' },
  {
    'data': '_snapshot/{repository}/{snapshot}/_restore',
    'value': '_snapshot/{repository}/{snapshot}/_restore'
  },
  { 'data': '_snapshot/_status', 'value': '_snapshot/_status' },
  {
    'data': '_snapshot/{repository}/_status',
    'value': '_snapshot/{repository}/_status'
  },
  {
    'data': '_snapshot/{repository}/{snapshot}/_status',
    'value': '_snapshot/{repository}/{snapshot}/_status'
  },
  { 'data': '_snapshot/{repository}/_verify', 'value': '_snapshot/{repository}/_verify' },
  {
    'data': '_tasks/_cancel',
    'value': '_tasks/_cancel'
  },
  { 'data': '_tasks/{task_id}/_cancel', 'value': '_tasks/{task_id}/_cancel' },
  {
    'data': '_tasks/{task_id}',
    'value': '_tasks/{task_id}'
  },
  { 'data': '_tasks', 'value': '_tasks' },
  {
    'data': '{indices}/{type}/_termvectors',
    'value': '{indices}/{type}/_termvectors'
  },
  {
    'data': '{indices}/{type}/{id}/_termvectors',
    'value': '{indices}/{type}/{id}/_termvectors'
  },
  {
    'data': '{indices}/_update_by_query',
    'value': '{indices}/_update_by_query'
  },
  {
    'data': '{indices}/{type}/_update_by_query',
    'value': '{indices}/{type}/_update_by_query'
  },
  {
    'data': '{indices}/{type}/{id}/_update',
    'value': '{indices}/{type}/{id}/_update'
  },
  {
    'data': '{indices}/_doc/{id}/_update',
    'value': '{indices}/_doc/{id}/_update'
  },
  {
    'data': '_ccr/auto_follow/{name}',
    'value': '_ccr/auto_follow/{name}'
  },
  { 'data': '{index}/_ccr/stats', 'value': '{index}/_ccr/stats' },
  {
    'data': '{index}/_ccr/follow',
    'value': '{index}/_ccr/follow'
  },
  { 'data': '_ccr/auto_follow', 'value': '_ccr/auto_follow' },
  {
    'data': '{index}/_ccr/pause_follow',
    'value': '{index}/_ccr/pause_follow'
  },
  { 'data': '{index}/_ccr/resume_follow', 'value': '{index}/_ccr/resume_follow' },
  {
    'data': '_ccr/stats',
    'value': '_ccr/stats'
  },
  { 'data': '{index}/_ccr/unfollow', 'value': '{index}/_ccr/unfollow' },
  {
    'data': '{indices}/_xpack/graph/_explore',
    'value': '{indices}/_xpack/graph/_explore'
  },
  {
    'data': '{indices}/{type}/_xpack/graph/_explore',
    'value': '{indices}/{type}/_xpack/graph/_explore'
  },
  { 'data': '_xpack', 'value': '_xpack' },
  {
    'data': '_xpack/license',
    'value': '_xpack/license'
  },
  {
    'data': '_xpack/license/basic_status',
    'value': '_xpack/license/basic_status'
  },
  {
    'data': '_xpack/license/trial_status',
    'value': '_xpack/license/trial_status'
  },
  { 'data': '_xpack/license/start_basic', 'value': '_xpack/license/start_basic' },
  {
    'data': '_xpack/license/start_trial',
    'value': '_xpack/license/start_trial'
  },
  {
    'data': '_xpack/migration/deprecations',
    'value': '_xpack/migration/deprecations'
  },
  {
    'data': '{indices}/_xpack/migration/deprecations',
    'value': '{indices}/_xpack/migration/deprecations'
  },
  {
    'data': '_xpack/migration/assistance',
    'value': '_xpack/migration/assistance'
  },
  {
    'data': '_xpack/migration/assistance/{indices}',
    'value': '_xpack/migration/assistance/{indices}'
  },
  {
    'data': '_xpack/migration/upgrade/{indices}',
    'value': '_xpack/migration/upgrade/{indices}'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_close',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_close'
  },
  {
    'data': '_xpack/ml/calendars/{calendar_id}/events/{event_id}',
    'value': '_xpack/ml/calendars/{calendar_id}/events/{event_id}'
  },
  {
    'data': '_xpack/ml/calendars/{calendar_id}/jobs/{job_id}',
    'value': '_xpack/ml/calendars/{calendar_id}/jobs/{job_id}'
  },
  {
    'data': '_xpack/ml/calendars/{calendar_id}',
    'value': '_xpack/ml/calendars/{calendar_id}'
  },
  {
    'data': '_xpack/ml/datafeeds/{datafeed_id}',
    'value': '_xpack/ml/datafeeds/{datafeed_id}'
  },
  {
    'data': '_xpack/ml/_delete_expired_data',
    'value': '_xpack/ml/_delete_expired_data'
  },
  {
    'data': '_xpack/ml/filters/{filter_id}',
    'value': '_xpack/ml/filters/{filter_id}'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}',
    'value': '_xpack/ml/anomaly_detectors/{job_id}'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots/{snapshot_id}',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots/{snapshot_id}'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_flush',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_flush'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_forecast',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_forecast'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/buckets/{timestamp}',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/buckets/{timestamp}'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/buckets',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/buckets'
  },
  {
    'data': '_xpack/ml/calendars/{calendar_id}/events',
    'value': '_xpack/ml/calendars/{calendar_id}/events'
  },
  {
    'data': '_xpack/ml/calendars',
    'value': '_xpack/ml/calendars'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/categories/{category_id}',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/categories/{category_id}'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/categories',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/categories'
  },
  {
    'data': '_xpack/ml/datafeeds/{datafeed_id}/_stats',
    'value': '_xpack/ml/datafeeds/{datafeed_id}/_stats'
  },
  { 'data': '_xpack/ml/datafeeds/_stats', 'value': '_xpack/ml/datafeeds/_stats' },
  {
    'data': '_xpack/ml/datafeeds',
    'value': '_xpack/ml/datafeeds'
  },
  {
    'data': '_xpack/ml/filters',
    'value': '_xpack/ml/filters'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/influencers',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/influencers'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/_stats',
    'value': '_xpack/ml/anomaly_detectors/_stats'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_stats',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_stats'
  },
  {
    'data': '_xpack/ml/anomaly_detectors',
    'value': '_xpack/ml/anomaly_detectors'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/overall_buckets',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/overall_buckets'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/results/records',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/results/records'
  },
  { 'data': '_xpack/ml/info', 'value': '_xpack/ml/info' },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_open',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_open'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_data',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_data'
  },
  {
    'data': '_xpack/ml/datafeeds/{datafeed_id}/_preview',
    'value': '_xpack/ml/datafeeds/{datafeed_id}/_preview'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots/{snapshot_id}/_revert',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots/{snapshot_id}/_revert'
  },
  {
    'data': '_xpack/ml/datafeeds/{datafeed_id}/_start',
    'value': '_xpack/ml/datafeeds/{datafeed_id}/_start'
  },
  {
    'data': '_xpack/ml/datafeeds/{datafeed_id}/_stop',
    'value': '_xpack/ml/datafeeds/{datafeed_id}/_stop'
  },
  {
    'data': '_xpack/ml/datafeeds/{datafeed_id}/_update',
    'value': '_xpack/ml/datafeeds/{datafeed_id}/_update'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/_update',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/_update'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots/{snapshot_id}/_update',
    'value': '_xpack/ml/anomaly_detectors/{job_id}/model_snapshots/{snapshot_id}/_update'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/_validate/detector',
    'value': '_xpack/ml/anomaly_detectors/_validate/detector'
  },
  {
    'data': '_xpack/ml/anomaly_detectors/_validate',
    'value': '_xpack/ml/anomaly_detectors/_validate'
  },
  { 'data': '_xpack/monitoring/_bulk', 'value': '_xpack/monitoring/_bulk' },
  {
    'data': '_xpack/monitoring/{type}/_bulk',
    'value': '_xpack/monitoring/{type}/_bulk'
  },
  { 'data': '_xpack/rollup/job/{id}', 'value': '_xpack/rollup/job/{id}' },
  {
    'data': '_xpack/rollup/job/',
    'value': '_xpack/rollup/job/'
  },
  { 'data': '_xpack/rollup/data/{id}', 'value': '_xpack/rollup/data/{id}' },
  {
    'data': '_xpack/rollup/data/',
    'value': '_xpack/rollup/data/'
  },
  {
    'data': '{indices}/_rollup_search',
    'value': '{indices}/_rollup_search'
  },
  {
    'data': '{indices}/{type}/_rollup_search',
    'value': '{indices}/{type}/_rollup_search'
  },
  {
    'data': '_xpack/rollup/job/{id}/_start',
    'value': '_xpack/rollup/job/{id}/_start'
  },
  {
    'data': '_xpack/rollup/job/{id}/_stop',
    'value': '_xpack/rollup/job/{id}/_stop'
  },
  {
    'data': '_xpack/security/_authenticate',
    'value': '_xpack/security/_authenticate'
  },
  {
    'data': '_xpack/security/user/{username}/_password',
    'value': '_xpack/security/user/{username}/_password'
  },
  {
    'data': '_xpack/security/user/_password',
    'value': '_xpack/security/user/_password'
  },
  {
    'data': '_xpack/security/realm/{realms}/_clear_cache',
    'value': '_xpack/security/realm/{realms}/_clear_cache'
  },
  {
    'data': '_xpack/security/role/{name}/_clear_cache',
    'value': '_xpack/security/role/{name}/_clear_cache'
  },
  {
    'data': '_xpack/security/role_mapping/{name}',
    'value': '_xpack/security/role_mapping/{name}'
  },
  {
    'data': '_xpack/security/role/{name}',
    'value': '_xpack/security/role/{name}'
  },
  {
    'data': '_xpack/security/user/{username}',
    'value': '_xpack/security/user/{username}'
  },
  {
    'data': '_xpack/security/user/{username}/_disable',
    'value': '_xpack/security/user/{username}/_disable'
  },
  {
    'data': '_xpack/security/user/{username}/_enable',
    'value': '_xpack/security/user/{username}/_enable'
  },
  { 'data': '_xpack/security/role_mapping', 'value': '_xpack/security/role_mapping' },
  {
    'data': '_xpack/security/role',
    'value': '_xpack/security/role'
  },
  { 'data': '_xpack/security/oauth2/token', 'value': '_xpack/security/oauth2/token' },
  {
    'data': '_xpack/security/user',
    'value': '_xpack/security/user'
  },
  { 'data': '_xpack/sql/close', 'value': '_xpack/sql/close' },
  {
    'data': '_xpack/sql',
    'value': '_xpack/sql'
  },
  { 'data': '_xpack/sql/translate', 'value': '_xpack/sql/translate' },
  {
    'data': '_xpack/ssl/certificates',
    'value': '_xpack/ssl/certificates'
  },
  { 'data': '_xpack/usage', 'value': '_xpack/usage' },
  {
    'data': '_xpack/watcher/watch/{watch_id}/_ack',
    'value': '_xpack/watcher/watch/{watch_id}/_ack'
  },
  {
    'data': '_xpack/watcher/watch/{watch_id}/_ack/{action_id}',
    'value': '_xpack/watcher/watch/{watch_id}/_ack/{action_id}'
  },
  {
    'data': '_xpack/watcher/watch/{watch_id}/_activate',
    'value': '_xpack/watcher/watch/{watch_id}/_activate'
  },
  {
    'data': '_xpack/watcher/watch/{watch_id}/_deactivate',
    'value': '_xpack/watcher/watch/{watch_id}/_deactivate'
  },
  {
    'data': '_xpack/watcher/watch/{id}',
    'value': '_xpack/watcher/watch/{id}'
  },
  {
    'data': '_xpack/watcher/watch/{id}/_execute',
    'value': '_xpack/watcher/watch/{id}/_execute'
  },
  {
    'data': '_xpack/watcher/watch/_execute',
    'value': '_xpack/watcher/watch/_execute'
  },
  { 'data': '_xpack/watcher/_start', 'value': '_xpack/watcher/_start' },
  {
    'data': '_xpack/watcher/stats',
    'value': '_xpack/watcher/stats'
  },
  {
    'data': '_xpack/watcher/stats/{metrics}',
    'value': '_xpack/watcher/stats/{metrics}'
  },
  { 'data': '_xpack/watcher/_stop', 'value': '_xpack/watcher/_stop' },
  {
    'data': '_processor',
    'value': '_processor'
  },
  { 'data': '_search/template/{id}', 'value': '_search/template/{id}' }]

export const esBodyKeyWords = [
  { caption: 'true', meta: 'true', value: `true` },
  { caption: 'false', meta: 'false', value: `false` },
  { caption: 'null', meta: 'null', value: `null` },
  { caption: '{}', meta: '{}', value: `{}` },
  { caption: '[]', meta: '[]', value: `[]` },
  { caption: 'aggs aggregations', meta: 'aggs', value: `
  "aggs": {
    "NAME": {
      "AGG_TYPE": {}
    }
  }
` },
  { caption: 'aliases', meta: 'aliases', value: `"aliases"` },
  { caption: 'filter', meta: 'filter', value: `"filter":{}` },
  { caption: 'post_filter', meta: 'post_filter', value: `"post_filter":{}` },

  { caption: 'bool', meta: 'bool', value: `"bool":{}` },
  { caption: 'exists', meta: 'exists', value: `"exists"` },
  { caption: 'query', meta: 'query', value: `"query":{
  
  }` },
  { caption: 'must', meta: 'must', value: `"must": [
     {}
  ], 
` },
  { caption: 'must_not', meta: 'must_not', value: `"must_not": [
     {}
  ], 
` },
  { caption: 'should', meta: 'should', value: `
  "should": [
     {}
  ], 
` },
  { caption: 'filter', meta: 'filter', value: `"filter"` },
  { caption: 'match', meta: 'match', value: `
  "match":{
    "key":"val"
  }
` },
  { caption: 'operator', meta: 'operator', value: `"operator"` },
  { caption: 'query_string', meta: 'query_string', value: `"query_string"` },
  { caption: 'range', meta: 'range', value: `
  "range":{
      "timeKey": {
      "from/to": "timeValue"
      }
  }
` },

  { caption: 'explain', meta: 'explain', value: `"explain": true` },
  { caption: 'gte >=', meta: 'gte', value: `"gte"` },
  { caption: 'gt >', meta: 'gt', value: `"gt"` },
  { caption: 'lte <=', meta: 'lte', value: `"lte"` },
  { caption: 'lt <', meta: 'lt', value: `"lt"` },
  { caption: 'boost', meta: 'boost', value: `"boost":1` },
  { caption: 'exists', meta: 'exists', value: `"exists"` },
  { caption: 'prefix', meta: 'prefix', value: `"
  "prefix": {
      "FIELD": {
      "value": ""
     }
  },   
"` },
  { caption: 'wildcard', meta: 'wildcard', value: `
  "wildcard": {
      "FIELD": {
        "value": "VALUE"
      }
    }
` },
  { caption: 'regexp', meta: 'regexp', value: `"
  "regexp": {
      "FIELD": "REGEXP"
    },   
"` },
  { caption: 'fuzzy', meta: 'fuzzy', value: `
  "fuzzy": {
    "FIELD": {}
  }  
` },
  { caption: 'type', meta: 'type', value: `"type":"typeVal"` },

  { caption: 'type long', meta: 'long', value: `"type":"long"` },
  { caption: 'type keyword', meta: 'keyword', value: `"type":"keyword"` },
  { caption: 'date', meta: 'date', value: `
       "type": "date",
       "format": "yyyy-MM-dd HH:mm:ss"
  ` },
  { caption: 'script update', meta: 'script', value: `"
  "script": {
    "source": "ctx._source['oldVal'] = 'newVal'"
  }
  "` },
  { caption: 'ctx._source.remove(key)', meta: 'ctx._source.remove(key)', value: `"ctx._source.remove(key)"` },
  { caption: 'terms', meta: 'terms', value: `"terms":" {
       "key": [ 
        "val1", 
        "val2", 
        "valN", 
      ] 
    }"` },
  { caption: 'text type', meta: 'text', value: `"type":"text"` },
  { caption: 'type keyword', meta: 'keyword', value: `"type":"keyword"` },
  { caption: 'type byte', meta: 'byte', value: `"type":"byte"` },
  { caption: 'type short', meta: 'short', value: `"type":"short"` },
  { caption: 'type integer', meta: 'integer', value: `"type":"integer"` },
  { caption: 'type long', meta: 'long', value: `"type":"long"` },
  { caption: 'type float', meta: 'float', value: `"type":"float"` },
  { caption: 'type half_float', meta: 'half_float', value: `"type":"half_float"` },
  { caption: 'type scaled_float', meta: 'scaled_float', value: `"type":"scaled_float"` },
  { caption: 'type double', meta: 'double', value: `"type":"double"` },
  { caption: 'type integer_range', meta: 'integer_range', value: `"type":"integer_range"` },
  { caption: 'type long_range', meta: 'long_range', value: `"type":"long_range"` },

  { caption: 'type float_range', meta: 'float_range', value: `"type":"float_range"` },
  { caption: 'type double_range', meta: 'double_range', value: `"type":"double_range"` },
  { caption: 'type date_range', meta: 'date_range', value: `"type":"date_range"` },
  { caption: 'type boolean', meta: 'boolean', value: `"type":"boolean"` },
  { caption: 'type binary', meta: 'binary', value: `"type":"binary"` },
  { caption: 'type object', meta: 'object', value: `"type":"object"` },
  { caption: 'type ip', meta: 'ip', value: `"type":"ip"` },

  { caption: 'type nested', meta: 'nested', value: `"type":"nested"` },
  { caption: 'query nested', meta: 'nested', value: `"nested": {
  }` },
  { caption: 'filter', meta: 'filter', value: `"filter":[]` },
  { caption: 'avg', meta: 'avg', value: `
  "avg" : {
      "field" : "key" 
   }` },
  { caption: 'min', meta: 'min', value: `
  "min" : {
      "field" : "min" 
   }` },
  { caption: 'max', meta: 'max', value: `
  "max" : {
      "field" : "key" 
   }` },
  { caption: 'sum', meta: 'sum', value: `
  "sum" : {
      "field" : "key" 
   }` },
  { caption: 'count', meta: 'count', value: `
  "count" : {
      "field" : "key" 
   }` },

  { caption: 'multi_match', meta: 'multi_match', value: `"multi_match ":{}` },

  { caption: 'ids', meta: 'ids', value: ` "ids": {}` },
  { caption: 'highlight', meta: 'highlight', value: `""highlight": {}"` },
  { caption: 'pre_tags', meta: 'pre_tags', value: `"pre_tags"` },
  { caption: 'post_tags', meta: 'post_tags', value: `"post_tags"` },
  { caption: 'fields', meta: 'fields', value: `"fields"` },
  { caption: 'highlight_query', meta: 'highlight_query', value: `"highlight_query:{}"` },
  { caption: 'fragment_size', meta: 'fragment_size', value: `"fragment_size":1` },
  { caption: 'number_of_fragments', meta: 'number_of_fragments', value: `"number_of_fragments":1` },
  { caption: 'no_match_size', meta: 'no_match_size', value: `"no_match_size":1` },

  { caption: 'from', meta: 'from', value: `"from":0` },
  { caption: 'size', meta: 'size', value: `"size":10` },
  { caption: 'sort', meta: 'sort', value: `"sort": [
    {
      "key": {
        "order": "asc/desc"
      }
    }
  ]
` },
  { caption: 'minimum_should_match', meta: 'minimum_should_match', value: `"minimum_should_match": 0` },

  { caption: 'format', meta: 'format', value: `"format": "dd/MM/yyyy||yyyy"` },
  { caption: 'field', meta: 'field', value: `"field":"key"` },
  { caption: 'size', meta: 'size', value: `"size":10` },
  { caption: '_source', meta: '_source', value: `
  "_source": [
  
  ]
` },
  { caption: 'collapse', meta: 'collapse', value: `
  "collapse": {
    "field": ""
  }
` },
  { caption: 'term', meta: 'term', value: `
  "term": {
    "key":"value"
  }
` },

  { caption: 'precision_threshold', meta: 'precision_threshold', value: `precision_threshold:40000` },
  { caption: 'cardinality', meta: 'cardinality', value: `
  "cardinality": 
  {
      "field": "key",
      "precision_threshold":40000
  }
` },

  { caption: 'dynamic', meta: 'dynamic', value: `"dynamic":false` },
  { caption: 'settings number_of_shards number_of_replicas', meta: 'settings', value: `
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 2
  }
` },
  { caption: 'properties', meta: 'properties', value: `
  "properties":{
  }
` },
  { caption: 'format', meta: 'format', value: `"format" : "yyyy-MM-dd HH:mm:ss"` },
  { caption: 'actions', meta: 'actions', value: `"actions":[]` },
  { caption: 'add', meta: 'add', value: `"add":{}` },
  { caption: 'index', meta: 'index', value: `"index":"indexName"` },
  { caption: 'alias', meta: 'alias', value: `"alias":"aliasName"` },

  { caption: 'refresh_interval', meta: 'refresh_interval', value: `"refresh_interval":"1s"` },

  { caption: 'number_of_shards', meta: 'number_of_shards', value: `"number_of_shards":1` },
  { caption: 'number_of_replicas', meta: 'number_of_replicas', value: `"number_of_replicas":1` },
  { caption: 'order', meta: 'order', value: `"order":"desc/asc"` },
  { caption: 'mappings', meta: 'mappings', value: `
  "mappings":{
        "_doc" : {
            "dynamic" : "false",
            "properties" : {}
        }
   }
` }
]

