<script lang="ts">
	import { conditionsFailed } from '../auth'
	import Router from 'svelte-spa-router'
	import routes from '../routes'
</script>

<body>
	<Router {routes} on:conditionsFailed={conditionsFailed} />
</body>

<style>

</style>