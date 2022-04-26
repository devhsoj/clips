<script lang="ts">
    import { link } from 'svelte-spa-router'

    let error = null

    async function signup(event) {
        try {
            const signupData = new URLSearchParams(new FormData(event.target)).toString()

            const res = await fetch('/auth/signup',{
                method:'POST',
                headers:{
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                body:signupData
            })

            if(res.status !== 200) {
                error = await res.text()
                return
            }

            window.location.href = '/'
        } catch(err) {
            console.trace(err)
        }
    }

</script>

<div class="content-wrapper">
    <section class="wrapper bg-soft-primary">
        <div class="container pt-5 pb-5 text-center">
            <div class="row">
                <div class="col-md-9 col-lg-7 col-xl-6 mx-auto">
                    <h2 class="display-1 mb-3">Project Clips</h2>
                </div>
            </div>
        </div>
    </section>

    <div class="container">
        <div class="row gx-md-11 gx-lg-0">
            <div id="signupContainer" class="col-xl-9 order-xl-2">
                <section id="signup" class="wrapper pt-8 pb-2">
                    <div class="col-lg-10 col-xl-8">
                        <form on:submit|preventDefault={signup}>
                            <h2 class="display-5 mb-7">Signup</h2>
                            <hr class="my-4" />

                            <div class="form-label-group mb-1">
                                <input id="username" name="username" class="form-control" placeholder="Username" required />
                                <label for="username">Username</label>
                            </div>

                            <div class="form-label-group mb-1">
                                <input id="password" name="password" class="form-control" placeholder="Password" type="password" required />
                                <label for="password">Password</label>
                            </div>

                            <div class="form-label-group mt-2">
                                <button type="submit" class="btn btn-primary btn-sm rounded-pill me-1 mb-4 mb-md-0 mt-2">Signup</button>
                            </div>

                            {#if error}
                                <div class="form-label-group mt-2">
                                    <span class="badge bg-red rounded-pill">
                                        {error}
                                    </span>
                                </div>
                            {/if}
                            <div class="form-label-group mt-2">
                                <a href="/login" use:link>Already have an account?</a>
                            </div>
                        </form>
                    </div>
                </section>
            </div>
            <aside class="col-xl-3 sidebar sticky-sidebar mt-md-0 py-16 d-none d-xl-block">
                <nav>
                    <ul class="nav list-unstyled lh-lg text-reset d-block">
                        <li class="nav-item"><a class="nav-link" href="/" use:link>Home</a></li>
                    </ul>
                </nav>
            </aside>
        </div>
    </div>
</div>

<style>

</style>