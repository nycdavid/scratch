package main

import (
	"context"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type (
	NoopReconciler struct{}
)

func (r *NoopReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// No-op: this is where your logic would go.
	// You can log to prove it’s wired up:
	ctrl.Log.Info("reconcile triggered", "namespacedName", req.NamespacedName)
	return ctrl.Result{}, nil
}

func (r *NoopReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.ConfigMap{}).
		Complete(r)
}

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))
}

func main() {
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))
	cfg := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:           scheme,
		LeaderElection:   false,
		LeaderElectionID: "noop-controller-example",
	})

	if err != nil {
		fmt.Println("[ERROR]: ", err)
		os.Exit(1)
	}

	if err := (&NoopReconciler{}).SetupWithManager(mgr); err != nil {
		fmt.Println("[ERROR]: unable to create controller: ", err)
		os.Exit(1)
	}

	fmt.Println("Starting manager...")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		fmt.Println("[ERROR]: problem running manager: ", err)
		os.Exit(1)
	}
}
