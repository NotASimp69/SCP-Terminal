# SCP Console

### Approach

1. Gather input prompt in SvelteKit
2. Send input to Django REST API
3. API call/web scraping to/from official "SCP Wiki"
   1. Outsource web scraping process onto async Celery worker
4. Send response to frontend via API
