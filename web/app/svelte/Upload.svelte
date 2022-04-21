<script lang="ts">
    let error

    async function uploadClip(event) {
        try {
            const clipData = new FormData(event.target)

            // No Content-Type header here so browser can work it's boundary magic!
            const res = await fetch('/api/clip/new',{
                method:'POST',
                body:clipData
            })

            if(res.status !== 200) {
                const data = await res.json()

                error = data.error

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
                        <form on:submit|preventDefault={uploadClip}>
                            <h2 class="display-5 mb-7">Upload Clip</h2>
                            <hr class="my-4" />

                            <div class="form-label-group mb-1">
                                <input id="title" name="title" class="form-control" placeholder="Title" required />
                                <label for="title">Title</label>
                            </div>

                            <div class="form-label-group mb-1">
                                <input id="description" name="description" class="form-control" placeholder="Description" required />
                                <label for="description">Description</label>
                            </div>

                            <div class="form-label-group mb-1">
                                <div>Select file to upload:</div>
                                <input type="file" name="clip" accept="video/*" required />
                            </div>

                            <div class="form-label-group mt-2">
                                <button type="submit" class="btn btn-primary btn-sm rounded-pill me-1 mb-4 mb-md-0 mt-2">Upload</button>
                            </div>

                            {#if error}
                                <div class="form-label-group mt-2">
                                    <span class="badge bg-red rounded-pill">
                                        {error}
                                    </span>
                                </div>
                            {/if}
                        </form>
                    </div>
                </section>
            </div>
            <aside class="col-xl-3 sidebar sticky-sidebar mt-md-0 py-16 d-none d-xl-block">
                <nav>
                    <ul class="nav list-unstyled lh-lg text-reset d-block">
                        <li class="nav-item"><a class="nav-link" href="#/">Home</a></li>
                        <li class="nav-item"><a class="nav-link" href="#/upload">Upload</a></li>
                    </ul>
                </nav>
            </aside>
        </div>
    </div>
</div>


<style>

</style>